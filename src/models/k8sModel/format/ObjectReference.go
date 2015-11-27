package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_objectreference
type ObjectReference struct {
	Kind string `json:"kind,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Name string `json:"name,omitempty"`
	// UID string `json:"uid,omitempty"`
	APIVersion string `json:"apiVersion,omitempty"`
	ResourceVersion string `json:"resourceVersion,omitempty"`
	FieldPath string `json:"fieldPath,omitempty"`
}
