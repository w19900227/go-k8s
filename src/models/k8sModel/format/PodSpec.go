package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_podspec
type PodSpec struct {
	Volumes []Volume `json:"volumes"`
	Containers []Container  `json:"containers"`
	RestartPolicy string `json:"restartPolicy"`
	TerminationGracePeriodSeconds int `json:"terminationGracePeriodSeconds"`
	// TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds"`
	ActiveDeadlineSeconds int `json:"activeDeadlineSeconds"`
	// ActiveDeadlineSeconds int64 `json:"activeDeadlineSeconds"`
	DnsPolicy string `json:"dnsPolicy"`
	NodeSelector map[string]string `json:"nodeSelector"` // any
	ServiceAccountName string `json:"serviceAccountName"`
	ServiceAccount string `json:"serviceAccount"`
	NodeName string `json:"nodeName"`
	HostNetwork bool `json:"hostNetwork"`
	ImagePullSecrets []LocalObjectReference `json:"imagePullSecrets"`
}
