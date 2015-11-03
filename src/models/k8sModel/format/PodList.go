package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podlist
type PodList struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ListMeta `json:"metadata"`
	Items []Pod `json:"items"`
}
