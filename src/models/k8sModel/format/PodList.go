package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podlist
type PodList struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	Items []Pod `json:"items,omitempty"`
}
