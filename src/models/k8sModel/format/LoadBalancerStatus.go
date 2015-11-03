package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_loadbalancerstatus
type LoadBalancerStatus struct {
	Ingress []LoadBalancerIngress `json:"ingress"`
}
