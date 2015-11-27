package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_endpointaddress
type EndpointAddress struct {
	IP string `json:"ip"`
	TargetRef ObjectReference `json:"targetRef"`
}
