package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_localobjectreference
type LocalObjectReference struct {
	Name string `json:"name"`
}
