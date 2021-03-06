package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_endpoints
type Endpoints struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ObjectMeta `json:"metadata"`
	Subsets []EndpointSubset `json:"subsets"`
}
