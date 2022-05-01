package models

const (
	StatusFetched   = "FETCHED"
	StatusCancelled = "CANCELLED"
)

type Entry struct {
	UID       string `json:"uid"`
	Timestamp int    `json:"timestamp"`
	VID       string `json:"vid"`
	State     string `json:"state"`
}
