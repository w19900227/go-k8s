package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_service
type Service struct {
	Kind string `json:"kind,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
	Metadata ObjectMeta `json:"metadata,omitempty"`
	Spec ServiceSpec `json:"spec,omitempty"`
	Status ServiceStatus `json:"status,omitempty"`
}
