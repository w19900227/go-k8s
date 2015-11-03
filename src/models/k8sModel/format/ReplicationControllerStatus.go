package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_replicationcontrollerstatus
type ReplicationControllerStatus struct {
	Replicas int `json:"replicas"`
	// Replicas int32 `json:"replicas"`
	ObservedGeneration int `json:"observedGeneration"`
	// ObservedGeneration int64 `json:"observedGeneration"`
}
