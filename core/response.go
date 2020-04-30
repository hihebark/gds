package core

// Response struct
type Response struct {
	Timestamp int64  `json:"timestamp"`
	Link      string `json:"link"`
	Status    int    `json:"status"`
	Length    string `json:"length"`
}
