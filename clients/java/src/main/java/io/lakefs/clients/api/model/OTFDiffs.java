/*
 * lakeFS API
 * lakeFS HTTP API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package io.lakefs.clients.api.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.lakefs.clients.api.model.DiffProperties;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * OTFDiffs
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class OTFDiffs {
  public static final String SERIALIZED_NAME_DIFFS = "diffs";
  @SerializedName(SERIALIZED_NAME_DIFFS)
  private List<DiffProperties> diffs = null;


  public OTFDiffs diffs(List<DiffProperties> diffs) {
    
    this.diffs = diffs;
    return this;
  }

  public OTFDiffs addDiffsItem(DiffProperties diffsItem) {
    if (this.diffs == null) {
      this.diffs = new ArrayList<DiffProperties>();
    }
    this.diffs.add(diffsItem);
    return this;
  }

   /**
   * Get diffs
   * @return diffs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<DiffProperties> getDiffs() {
    return diffs;
  }


  public void setDiffs(List<DiffProperties> diffs) {
    this.diffs = diffs;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    OTFDiffs otFDiffs = (OTFDiffs) o;
    return Objects.equals(this.diffs, otFDiffs.diffs);
  }

  @Override
  public int hashCode() {
    return Objects.hash(diffs);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class OTFDiffs {\n");
    sb.append("    diffs: ").append(toIndentedString(diffs)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

