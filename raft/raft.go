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
	RequestVote       chan *Node
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
		RequestVote:       make(chan *Node),
		VoteChan:          make(chan int),
		AppendEntriesChan: make(chan int),
	}
}

func (n *Node) Active(otherNodes []*Node) {
	ElectionTimeout := n.ElectionTimeout

	go func() {
		for {
			if n.Rule == Follower {
				time.Sleep(ElectionTimeout)
			} else if n.Rule == Candidate {
				for _, node := range otherNodes {
					node.RequestVote <- n
				}
			} else {
				for _, node := range otherNodes {
					node.AppendEntriesChan <- n.ID
				}
			}
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
			if n.Rule == Follower {
				n.Voted++
				if (len(otherNodes) / 2) <= n.Voted {
					n.LeaderID = n.ID
					n.Rule = Leader
					n.Voted = 0
				}
			}
		case node := <-n.RequestVote:
			node.VoteChan <- n.ID
		}
	}()
}

func StartElection(nodes []*Node) {
	for i, node := range nodes {
		go node.Active(append(nodes[:i], nodes[i+1:]...))
	}
}
