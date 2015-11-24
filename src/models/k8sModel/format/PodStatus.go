package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podstatus
type PodStatus struct {
	Phase string `json:"phase,omitempty"`
	Conditions []PodCondition `json:"conditions,omitempty"`
	Message string `json:"message,omitempty"`
	Reason string `json:"reason,omitempty"`
	HostIP string `json:"hostIP,omitempty"`
	PodIP string `json:"podIP,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	ContainerStatuses []ContainerStatus `json:"containerStatuses,omitempty"`
}
