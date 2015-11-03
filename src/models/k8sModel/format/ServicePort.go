package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_serviceport
type ServicePort struct {
	Name string `json:"name"`
	Protocol string `json:"protocol"`
	// Port int32 `json:"port"`
	Port int `json:"port"`
	TargetPort int `json:"targetPort"`
	// NodePort int32 `json:"nodePort"`
	NodePort int `json:"nodePort"`
}
