package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_servicespec
type ServiceSpec struct {
	Ports []ServicePort `json:"ports,omitempty"`
	Selector map[string]string `json:"selector,omitempty"` // any
	ClusterIP string `json:"clusterIP,omitempty"`
	Type string `json:"type,omitempty"`
	DeprecatedPublicIPs []string `json:"deprecatedPublicIPs,omitempty"` //string array
	SessionAffinity string `json:"sessionAffinity,omitempty"`
}
