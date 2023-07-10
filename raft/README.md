# raft
The Raft consensus algorithm is a protocol designed to ensure fault-tolerance and leader election in distributed systems.

In Raft, a cluster of nodes elects a leader among them to coordinate their activities. The leader is responsible for handling client requests, replicating data to other nodes, and maintaining consistency across the cluster. When the leader fails or becomes unavailable, a new leader needs to be elected through the Raft election algorithm.
Each node in the system has its own internal clock. Nodes randomly choose a timeout value, let's say between 150 and 300 milliseconds.
During normal operation, a node sends out "heartbeat" messages to other nodes to show that it is still active and functioning properly.
If a node doesn't receive any heartbeat message from the leader within its timeout period, it assumes that something may have gone wrong with the leader.
The node that suspects a problem starts an election by sending a special message called a "RequestVote" to all other nodes in the system. 
When a node receives a RequestVote, it checks whether it has already voted in this term. If it hasn't voted for anyone yet, and it agrees that the requesting node is eligible to become the leader (based on the requesting node's log information), it grants its vote by sending a "Vote" message back.
If a candidate receives votes from the majority of the nodes, it becomes the new leader for the next term. It notifies everyone about its victory by sending "AppendEntries" messages, indicating that it is the new leader and others should follow it.
Once a node receives the AppendEntries message from a valid leader, it resets its timer and recognizes the new leader as the authority.
The leader continues to send periodic heartbeat messages to maintain its leadership and keep the system running smoothly.