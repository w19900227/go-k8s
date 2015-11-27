package format

// http://kubernetes.io/v1.0/docs/api-reference/definitions.html#_v1_eventsource
type EventSource struct {
	Component string `json:"component,omitempty"`
	Host string `json:"host,omitempty"`
}
