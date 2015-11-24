package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_pod
type Pod struct {
	Kind string `json:"kind,omitempty"`
	ApiVersion string `json:"apiVersion,omitempty"`
	Metadata ObjectMeta `json:"metadata,omitempty"`
	Spec PodSpec `json:"spec,omitempty"`
	Status PodStatus `json:"status,omitempty"`
}
