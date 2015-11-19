package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_objectmeta
type ObjectMeta struct {
	Name string `json:"name,omitempty"`
    GenerateName string `json:"generateName,omitempty"`
    Namespace string `json:"namespace,omitempty"`
    SelfLink string `json:"selfLink,omitempty"`
    Uid string `json:"uid,omitempty"`
    ResourceVersion string `json:"resourceVersion,omitempty"`
    Generation int `json:"generation,omitempty"`
    // Generation int64 `json:"generation,omitempty"`
    CreationTimestamp string `json:"creationTimestamp,omitempty"`
    DeletionTimestamp string `json:"deletionTimestamp,omitempty"`
    Labels map[string]string `json:"labels,omitempty"` //any
    // Annotations {} any `json:"annotations,omitempty"` //any
}
