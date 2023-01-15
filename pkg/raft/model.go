package raft

type AppendEntry struct {
	TimeStamp int64  `json:"timeStamp"`
	Payload   string `json:"payload"`
}
