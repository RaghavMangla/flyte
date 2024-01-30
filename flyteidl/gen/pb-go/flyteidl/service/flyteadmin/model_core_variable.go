/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Defines a strongly typed variable.
type CoreVariable struct {
	// Variable literal type.
	Type_ *CoreLiteralType `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	// +optional This object allows the user to specify how Artifacts are created. name, tag, partitions can be specified. The other fields (version and project/domain) are ignored.
	ArtifactPartialId *CoreArtifactId `json:"artifact_partial_id,omitempty"`
	ArtifactTag *CoreArtifactTag `json:"artifact_tag,omitempty"`
}
