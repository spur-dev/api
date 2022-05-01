package models

const (
	StatusProcessing = "PROCESSING"
	StatusFailed     = "FAILED"
	StatusReady      = "READY"
	StatusDelete     = "DELETE"
)

type MetaData struct { // TODO : add time stamp
	VID       string `json:"id"`
	Timestamp int    `json:"timestamp"`
	UID       string `json:"uid,omitempty"`
	State     string `json:"status,omitempty"` // TODO: rename to State
}
