package raft

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

type Fsm struct {
	Name               string
	Nodes              []string
	CurrentLeader      string
	State              NodeState
	LastHeartbeat      int64
	ElectionTerm       int
	ElectionTimeout    int64
	HeartbeatTimeout   int
	HeartbeatsReceiver <-chan int64
}

func NewFsm(
	name string,
	nodes []string,
	heartbeatsReceiver <-chan int64,
) *Fsm {
	return &Fsm{
		Name:               name,
		Nodes:              nodes,
		State:              Follower,
		ElectionTerm:       0,
		ElectionTimeout:    int64(rand.Intn(150) + 5050),
		HeartbeatTimeout:   50,
		HeartbeatsReceiver: heartbeatsReceiver,
	}
}

func (r *Fsm) Run() {
	for {
		if r.State == Follower {
			r.logCurrentState()
			expired := false

			for expired != true {
				select {
				case hb := <-r.HeartbeatsReceiver:
					log.Debugf("Got new heartbeat: %d", hb)

					timeSinceLastHeartbeat := hb - r.LastHeartbeat
					if timeSinceLastHeartbeat > r.ElectionTimeout {
						log.Warnf("Heartbeat was expired for %dms", timeSinceLastHeartbeat)
						r.becomeCandidate()
						expired = true
						break
					} else {
						r.LastHeartbeat = hb
						// respond to master???
					}
				case <-time.After(time.Duration(r.ElectionTimeout) * time.Millisecond):
					fmt.Println("Timeout: News feed finished")
					r.becomeCandidate()
					expired = true
					break
				}
			}
		} else if r.State == Candidate {
			// todo
			r.logCurrentState()
			hb := <-r.HeartbeatsReceiver
			if hb > r.LastHeartbeat {
				r.LastHeartbeat = hb
				r.State = Follower
			}
		} else {
			r.logCurrentState()
		}
	}
}

func (r *Fsm) logCurrentState() {
	log.Debugf("Node %s's state is %s", r.Name, r.State.String())
}

func (r *Fsm) becomeCandidate() {
	r.State = Candidate
}
