package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_serviceport
type ServicePort struct {
	Name string `json:"name,omitempty"`
	Protocol string `json:"protocol,omitempty"`
	// Port int32 `json:"port,omitempty"`
	Port int `json:"port,omitempty"`
	TargetPort int `json:"targetPort,omitempty"`
	// NodePort int32 `json:"nodePort,omitempty"`
	NodePort int `json:"nodePort,omitempty"`
}
