package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_replicationcontrollerlist
type ReplicationControllerList struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ListMeta `json:"metadata"`
	Items []ReplicationController `json:"items"`
}
