package bully

type Role int

const (
	Follower Role = iota
	Candidate
	Leader
)

type Node struct {
	ID       int
	LeaderID int
	Role     Role
}
