package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_nodespec
type NodeSpec struct {
	PodCIDR string `json:"podCIDR,omitempty"`
	ExternalID string `json:"externalID,omitempty"`
	ProviderID string `json:"providerID,omitempty"`
	Unschedulable bool `json:"unschedulable,omitempty"`
}
