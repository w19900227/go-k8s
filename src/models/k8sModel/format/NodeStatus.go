package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodestatus
type NodeStatus struct {
	Capacity ResourceList `json:"capacity,omitempty"` //any
	Phase string `json:"phase,omitempty"`
	Conditions []NodeCondition `json:"conditions,omitempty"`
	Addresses []NodeAddress `json:"addresses,omitempty"`
	NodeInfo NodeSystemInfo `json:"nodeInfo,omitempty"`
	// DaemonEndpoints NodeDaemonEndpoints `json:"daemonEndpoints,omitempty"`
}

type ResourceList map[string]string
// type ResourceName string
// type ResourceList map[ResourceName]resource.Quantity
