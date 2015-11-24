package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podspec
type PodSpec struct {
	Volumes []Volume `json:"volumes,omitempty"`
	Containers []Container  `json:"containers,omitempty"`
	RestartPolicy string `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds int `json:"terminationGracePeriodSeconds,omitempty"`
	// TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds int `json:"activeDeadlineSeconds,omitempty"`
	// ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds,omitempty"`
	DnsPolicy string `json:"dnsPolicy,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"` // any
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	ServiceAccount string `json:"serviceAccount,omitempty"`
	NodeName string `json:"nodeName,omitempty"`
	HostNetwork bool `json:"hostNetwork,omitempty"`
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets,omitempty"`
}
