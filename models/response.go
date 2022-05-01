package models

type NewVideoResponse struct {
	VID       string `json:"vid"`
	UID       string `json:"uid"`
	Timestamp int    `json:"timestamp"`
}

type MsgResponse struct {
	Msg string `json:"msg"`
}

type GetVideoResponse struct {
	VID     string `json:"vid"`
	UID     string `json:"uid"`
	Preview string `json:"preview"`
	Src     string `json:"src,omitempty"`
	State   string `json:"state,omitempty"`
}

type LambdaResponse struct {
	Status string `json:"status"`
	// VideoName string `json:"vname"`
	Msg string `json:"msg,omitempty"`
}
