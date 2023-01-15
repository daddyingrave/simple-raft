package raft

type NodeState int32

const (
	Follower NodeState = iota
	Candidate
	Leader
)

func (r NodeState) String() string {
	switch r {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	}
	panic("impossible")
}
