# Bully

The **Bully algorithm** is a classic algorithm used in distributed systems to elect a leader or coordinator among a group of processes. The purpose of the algorithm is to ensure that only one process becomes the coordinator, even in situations where multiple processes may attempt to become the leader simultaneously.

## Process Identification
Each process is given a unique ID based on its importance.

## Election Process
1. When a process notices that the current leader has failed, it starts an election by sending messages to processes with higher IDs.
2. Processes receiving an election message decide whether to participate or not. If they have a higher ID and are not already in an election, they respond with an OK message.
3. If no higher-ID process responds within a certain time, the initiating process becomes the leader.
4. If a higher-ID process responds with an OK message, the initiating process withdraws from the election.
5. Once the initiating process receives OK messages from all higher-ID processes or the timeout expires, it broadcasts a victory message to declare itself as the new leader.

## Updating State
Other processes update their state upon receiving the victory message.

The Bully algorithm ensures that the process with the highest ID becomes the leader. However, it assumes reliable communication and may not be ideal for complex distributed systems with frequent failures.