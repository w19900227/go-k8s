package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodecondition
type NodeCondition struct {
	Types string `json:"types,omitempty"`
	Status string `json:"status,omitempty"`
	LastHeartbeatTime string `json:"lastHeartbeatTime,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	Reason string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}
