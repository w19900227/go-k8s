package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_pod
type Pod struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ObjectMeta `json:"metadata"`
	Spec PodSpec `json:"spec"`
	Status PodStatus `json:"status"`
}
