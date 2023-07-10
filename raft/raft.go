package raft

import "time"

type Rule int

const (
	Follower Rule = iota
	Candidate
	Leader
)

type Node struct {
	ID                int
	LeaderID          int
	ElectionTimeout   time.Duration
	HeartBeatInterval time.Duration
	Rule              Rule
	VoteChan          chan string
}

func NewNode(id, electionTimeout, heartbeatInterval int) *Node {
	return &Node{
		ID:                id,
		LeaderID:          -1,
		ElectionTimeout:   time.Duration(electionTimeout) * time.Millisecond,
		HeartBeatInterval: time.Duration(heartbeatInterval) * time.Millisecond,
		Rule:              Follower,
	}
}

func (n *Node) Active() {

}

func StartElection(nodes []*Node) {
	for _, node := range nodes {
		go node.Active()
	}
}
