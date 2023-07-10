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
	Voted             int
	VoteChan          chan int
	AppendEntriesChan chan int
}

func NewNode(id, electionTimeout, heartbeatInterval int) *Node {
	return &Node{
		ID:                id,
		LeaderID:          -1,
		ElectionTimeout:   time.Duration(electionTimeout) * time.Millisecond,
		HeartBeatInterval: time.Duration(heartbeatInterval) * time.Millisecond,
		Voted:             0,
		Rule:              Follower,
		VoteChan:          make(chan int),
		AppendEntriesChan: make(chan int),
	}
}

func (n *Node) Active(count int) {
	ElectionTimeout := n.ElectionTimeout

	go func() {
		for {
			time.Sleep(ElectionTimeout)
		}
	}()

	go func() {
		select {
		case l := <-n.AppendEntriesChan:
			n.LeaderID = l
			n.Voted = 0
			n.Rule = Follower
			ElectionTimeout = n.ElectionTimeout
		case <-n.VoteChan:
			n.Voted++
			if (count / 2) <= n.Voted {
				n.LeaderID = n.ID
				n.Rule = Leader
				n.Voted = 0
			}
		}
	}()
}

func StartElection(nodes []*Node) {
	for _, node := range nodes {
		go node.Active(len(nodes))
	}
}
