package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podstatus
type PodStatus struct {
	Phase string `json:"phase"`
	Conditions []PodCondition `json:"conditions"`
	Message string `json:"message"`
	Reason string `json:"reason"`
	HostIP string `json:"hostIP"`
	PodIP string `json:"podIP"`
	StartTime string `json:"startTime"`
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
}
