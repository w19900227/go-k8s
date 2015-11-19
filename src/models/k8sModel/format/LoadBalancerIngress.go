package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_loadbalanceringress
type LoadBalancerIngress struct {
	Ip string `json:"ip,omitempty"`
	Hostname string `json:"hostname,omitempty"`
}
