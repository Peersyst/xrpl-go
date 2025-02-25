package types

// The consensus stream sends consensusPhase messages when the consensus process changes phase.
// The message contains the new phase of consensus the server is in.
type ConsensusStream struct {
	// The value `consensusPhase` indicates this is from the consensus stream
	Type Type `json:"type"`
	// The new consensus phase the server is in. Possible values are `open`, `establish`, and `accepted`.
	Consensus string `json:"consensus"`
}
