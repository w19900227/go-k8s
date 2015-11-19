package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_servicestatus
type ServiceStatus struct {
	LoadBalancer LoadBalancerStatus `json:"loadBalancer,omitempty"`
}
