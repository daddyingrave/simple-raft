package raft

import (
	log "github.com/sirupsen/logrus"
)

type Service interface {
	Acknowledge(entry AppendEntry)
}

type service struct {
	heartbeatsReceiver chan<- int64
}

func NewService(heartbeatsReceiver chan<- int64) Service {
	return &service{heartbeatsReceiver: heartbeatsReceiver}
}

func (r *service) Acknowledge(entry AppendEntry) {
	log.Debugf("Got AppendEntry with timestamp: %d and payload: %s", entry.TimeStamp, entry.Payload)

	r.heartbeatsReceiver <- entry.TimeStamp
}
