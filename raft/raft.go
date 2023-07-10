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
