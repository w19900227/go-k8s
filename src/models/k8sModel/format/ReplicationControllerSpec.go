package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_replicationcontrollerspec
type ReplicationControllerSpec struct {
	Replicas int `json:"replicas"`
	// Replicas int32 `json:"replicas"`
	Selector  map[string]string `json:"selector"` //any
	Template PodTemplateSpec `json:"template"`
}
