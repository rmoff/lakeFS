package nessie

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/treeverse/lakefs/pkg/api"

	"github.com/stretchr/testify/require"
)

func TestCommitSingle(t *testing.T) {
	for _, direct := range []bool{false, true} {
		name := "indirect"
		if direct {
			name = "direct"
		}
		t.Run(name, func(t *testing.T) {
			ctx, _, repo := setupTest(t)
			objPath := "1.txt"

			_, objContent := uploadFileRandomData(ctx, t, repo, masterBranch, objPath, direct)
			commitResp, err := client.CommitWithResponse(ctx, repo, masterBranch, api.CommitJSONRequestBody{
				Message: "nessie:singleCommit",
			})
			require.NoError(t, err, "failed to commit changes")
			require.Equal(t, http.StatusCreated, commitResp.StatusCode())

			getObjResp, err := client.GetObjectWithResponse(ctx, repo, masterBranch, &api.GetObjectParams{Path: objPath})
			require.NoError(t, err, "failed to get object")
			require.Equal(t, http.StatusOK, getObjResp.StatusCode())

			body := string(getObjResp.Body)
			require.Equal(t, objContent, body, fmt.Sprintf("path: %s, expected: %s, actual:%s", objPath, objContent, body))
		})
	}
}

// genNames generates n consecutive filenames starting with prefix.
func genNames(n int, prefix string) []string {
	ret := make([]string, n)
	for i := range ret {
		ret[i] = fmt.Sprint(prefix, i+1)
	}
	return ret
}

type Upload struct {
	Repo, Branch, Path string
}

// upload uploads random file data for uploads.
func upload(ctx context.Context, uploads chan Upload, direct bool) error {
	for u := range uploads {
		_, _, err := uploadFileRandomDataAndReport(ctx, u.Repo, u.Branch, u.Path, direct)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestCommitInMixedOrder(t *testing.T) {
	const (
		parallelism = 5
		size        = 100
	)

	for _, direct := range []bool{false, true} {
		name := "indirect"
		if direct {
			name = "direct"
		}
		t.Run(name, func(t *testing.T) {
			ctx, _, repo := setupTest(t)

			names1 := genNames(size, "run2/foo")
			uploads := make(chan Upload, size)
			wg := sync.WaitGroup{}
			for i := 0; i < parallelism; i++ {
				wg.Add(1)
				go func() {
					if err := upload(ctx, uploads, direct); err != nil {
						t.Error(err)
					}
					wg.Done()
				}()
			}
			for _, name := range names1 {
				uploads <- Upload{Repo: repo, Branch: masterBranch, Path: name}
			}
			close(uploads)
			wg.Wait()

			if t.Failed() {
				t.FailNow()
			}

			commitResp, err := client.CommitWithResponse(ctx, repo, masterBranch, api.CommitJSONRequestBody{
				Message: "nessie:mixedOrderCommit1",
			})
			require.NoError(t, err, "failed to commit changes")
			require.Equal(t, http.StatusCreated, commitResp.HTTPResponse.StatusCode)

			names2 := genNames(size, "run1/foo")
			uploads = make(chan Upload, size)
			wg = sync.WaitGroup{}
			for i := 0; i < parallelism; i++ {
				wg.Add(1)
				go func() {
					if err := upload(ctx, uploads, direct); err != nil {
						t.Error(err)
					}
					wg.Done()
				}()
			}
			for _, name := range names2 {
				uploads <- Upload{Repo: repo, Branch: masterBranch, Path: name}
			}
			close(uploads)
			wg.Wait()

			if t.Failed() {
				t.FailNow()
			}

			commitResp, err = client.CommitWithResponse(ctx, repo, masterBranch, api.CommitJSONRequestBody{
				Message: "nessie:mixedOrderCommit2",
			})
			require.NoError(t, err, "failed to commit second set of changes")
			require.Equal(t, http.StatusCreated, commitResp.HTTPResponse.StatusCode)
		})
	}
}
