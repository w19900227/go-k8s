package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_containerstatus
type ContainerStatus struct {
	Name string `json:"name,omitempty"`
	State ContainerState `json:"state,omitempty"`
	LastState ContainerState `json:"lastState,omitempty"`
	Ready bool `json:"ready,omitempty"`
	// RestartCount int32 `json:"restartCount,omitempty"`
	RestartCount int `json:"restartCount,omitempty"`
	Image string `json:"image,omitempty"`
	ImageID string `json:"imageID,omitempty"`
	ContainerID string `json:"containerID,omitempty"`
}
