package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
	"github.com/treeverse/lakefs/pkg/api"
	"github.com/treeverse/lakefs/pkg/fileutil"
	"github.com/treeverse/lakefs/pkg/local"
	"github.com/treeverse/lakefs/pkg/uri"
)

const (
	localCloneMinArgs = 1
	localCloneMaxArgs = 2
	filesChanSize     = 1000
)

var localCloneCmd = &cobra.Command{
	Use:   "clone <path uri> [directory]",
	Short: "Clone a path from a lakeFS repository into a new directory.",
	Args:  cobra.RangeArgs(localCloneMinArgs, localCloneMaxArgs),
	Run: func(cmd *cobra.Command, args []string) {
		client := getClient()
		remote, localPath := getLocalArgs(args, true, false)
		syncFlags := getLocalSyncFlags(cmd, client)
		updateIgnore := Must(cmd.Flags().GetBool(localGitIgnoreFlagName))
		empty, err := fileutil.IsDirEmpty(localPath)
		if err != nil {
			if !errors.Is(err, fs.ErrNotExist) {
				DieErr(err)
			}
		} else if !empty {
			DieFmt("directory '%s' exists and is not empty", localPath)
		}

		head, err := localInit(cmd.Context(), localPath, remote, false, updateIgnore)
		if err != nil {
			DieErr(err)
		}
		stableRemote := remote.WithRef(head)
		// Dynamically construct changes
		c := make(chan *local.Change, filesChanSize)
		go func() {
			defer close(c)
			hasMore := true
			var after string
			for hasMore {
				listResp, err := client.ListObjectsWithResponse(cmd.Context(), remote.Repository, stableRemote.Ref, &api.ListObjectsParams{
					After:        (*api.PaginationAfter)(swag.String(after)),
					Prefix:       (*api.PaginationPrefix)(remote.Path),
					UserMetadata: swag.Bool(true),
				})
				DieOnErrorOrUnexpectedStatusCode(listResp, err, http.StatusOK)

				for _, o := range listResp.JSON200.Results {
					path := strings.TrimPrefix(o.Path, remote.GetPath())
					// skip directory markers
					if path == "" || (strings.HasSuffix(path, uri.PathSeparator) && swag.Int64Value(o.SizeBytes) == 0) {
						continue
					}
					path = strings.TrimPrefix(path, uri.PathSeparator)
					c <- &local.Change{
						Source: local.ChangeSourceRemote,
						Path:   path,
						Type:   local.ChangeTypeAdded,
					}
				}
				hasMore = listResp.JSON200.Pagination.HasMore
				after = listResp.JSON200.Pagination.NextOffset
			}
		}()
		sigCtx := localHandleSyncInterrupt(cmd.Context())
		s := local.NewSyncManager(sigCtx, client, syncFlags.parallelism, syncFlags.presign)
		err = s.Sync(localPath, stableRemote, c)
		if err != nil {
			DieErr(err)
		}
		fmt.Printf("\nSuccessfully cloned %s to %s.\n", remote, localPath)
		Write(localSummaryTemplate, struct {
			Operation string
			local.Tasks
		}{
			Operation: "Clone",
			Tasks:     s.Summary(),
		})
	},
}

//nolint:gochecknoinits
func init() {
	withGitIgnoreFlag(localCloneCmd)
	withLocalSyncFlags(localCloneCmd)
	localCmd.AddCommand(localCloneCmd)
}
