package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_objectmeta
type ObjectMeta struct {
	Name string `json:"name"`
    GenerateName string `json:"generateName"`
    Namespace string `json:"namespace"`
    SelfLink string `json:"selfLink"`
    Uid string `json:"uid"`
    ResourceVersion string `json:"resourceVersion"`
    Generation int `json:"generation"`
    // Generation int64 `json:"generation"`
    CreationTimestamp string `json:"creationTimestamp"`
    DeletionTimestamp string `json:"deletionTimestamp"`
    Labels map[string]string `json:"labels"` //any
    // Annotations {} any `json:"annotations"` //any
}
