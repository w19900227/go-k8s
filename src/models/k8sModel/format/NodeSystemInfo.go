package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodesysteminfo
type NodeSystemInfo struct {
	MachineID string `json:"machineID,omitempty"`
	SystemUUID string `json:"systemUUID,omitempty"`
	BootID string `json:"bootID,omitempty"`
	KernelVersion string `json:"kernelVersion,omitempty"`
	OsImage string `json:"osImage,omitempty"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion,omitempty"`
	KubeletVersion string `json:"kubeletVersion,omitempty"`
	KubeProxyVersion string `json:"kubeProxyVersion,omitempty"`
}
