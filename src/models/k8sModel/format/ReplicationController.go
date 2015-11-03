package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_replicationcontroller
type ReplicationController struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ObjectMeta `json:"metadata"`
	Spec ReplicationControllerSpec `json:"spec"`
	Status ReplicationControllerStatus `json:"status"`
}
