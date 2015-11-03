package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podtemplatespec
type PodTemplateSpec struct {
	Metadata ObjectMeta `json:"metadata"`
	Spec PodSpec `json:"spec"`
}
