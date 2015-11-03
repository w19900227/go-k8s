package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_service
type Service struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata ObjectMeta `json:"metadata"`
	Spec ServiceSpec `json:"spec"`
	Status ServiceStatus `json:"status"`
}
