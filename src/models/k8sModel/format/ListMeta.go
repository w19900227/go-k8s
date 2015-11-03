package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_listmeta
type ListMeta struct {
	SelfLink string `json:"selfLink"`
	ResourceVersion string `json:"	resourceVersion"`
}
