package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_containerport
type ContainerPort struct {
	Name string `json:"name"`
	HostPort int `json:"hostPort"`
	// HostPort int32 `json:"hostPort"`
	ContainerPort int `json:"containerPort"`
	// ContainerPort int32 `json:"containerPort"`
	Protocol string `json:"protocol"`
	HostIP string `json:"hostIP"`
}
