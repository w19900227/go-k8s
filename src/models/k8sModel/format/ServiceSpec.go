package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_servicespec
type ServiceSpec struct {
	Ports []ServicePort `json:"ports"`
	Selector map[string]string `json:"selector"` // any
	ClusterIP string `json:"clusterIP"`
	Types string `json:"types"`
	DeprecatedPublicIPs []string `json:"deprecatedPublicIPs"` //string array
	SessionAffinity string `json:"sessionAffinity"`
}
