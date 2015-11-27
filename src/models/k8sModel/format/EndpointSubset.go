package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_endpointsubset
type EndpointSubset struct {
	Addresses []EndpointAddress `json:"addresses"`
	// NotReadyAddresses []EndpointAddress `json:"NotReadyAddresses"`
	Ports []EndpointPort `json:"ports"`
}
