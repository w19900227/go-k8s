package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodeaddress
type NodeAddress struct {
	Types string `json:"types,omitempty"`
	Address string `json:"address,omitempty"`
}
