package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_event
type Event struct {
	// Kind string `json:"kind"`
	// ApiVersion string `json:"apiVersion"`
	Metadata ObjectMeta `json:"metadata,omitempty"`
	InvolvedObject ObjectReference `json:"involvedObject,omitempty"`
	Reason string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	Source EventSource `json:"source,omitempty"`
	FirstTimestamp string `json:"firstTimestamp,omitempty"`
	LastTimestamp string `json:"lastTimestamp,omitempty"`
	Count int `json:"count,omitempty"`
}
