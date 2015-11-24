package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodelist
type NodeList struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	Items []Node `json:"items,omitempty"`
}
