package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podcondition
type PodCondition struct {
	Types string `json:"types,omitempty"`
	Status string `json:"status,omitempty"`
}
