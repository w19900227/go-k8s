package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_containerstate
type ContainerState struct {
	// Waiting ContainerStateWaiting `json:"waiting"`
	// Running ContainerStateRunning `json:"running"`
	// Terminated ContainerStateTerminated `json:"terminated"`
}
