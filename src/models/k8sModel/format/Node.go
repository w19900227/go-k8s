package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_node
type Node struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Metadata ObjectMeta `json:"metadata,omitempty"`
	Spec NodeSpec `json:"spec,omitempty"`
	Status NodeStatus `json:"status,omitempty"`
}
