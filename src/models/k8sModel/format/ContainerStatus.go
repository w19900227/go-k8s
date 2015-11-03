package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_containerstatus
type ContainerStatus struct {
	Name string `json:"name"`
	State ContainerState `json:"state"`
	LastState ContainerState `json:"lastState"`
	Ready bool `json:"ready"`
	// RestartCount int32 `json:"restartCount"`
	RestartCount int `json:"restartCount"`
	Image string `json:"image"`
	ImageID string `json:"imageID"`
	ContainerID string `json:"containerID"`
}
