package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_container
type Container struct {
	Name string `json:"name"`
	Image string `json:"image"`
	// Command map[string]string `json:"command"` //string array
	// Args map[string]string `json:"args"` //string array
	// WorkingDir string `json:"workingDir"`
	Ports []ContainerPort `json:"ports"`
	// Env []EnvVar `json:"env"`
	// Resources ResourceRequirements `json:"resources"`
	// VolumeMounts []VolumeMount `json:"volumeMounts"`
	// LivenessProbe Probe `json:"livenessProbe"`
	// ReadinessProbe Probe `json:"readinessProbe"`
	// Lifecycle Lifecycle `json:"lifecycle"`
	TerminationMessagePath string `json:"terminationMessagePath"`
	ImagePullPolicy string `json:"imagePullPolicy"`
	// SecurityContext SecurityContext `json:"securityContext"`
}
