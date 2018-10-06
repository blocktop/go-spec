package spec

// Competition represents the latest block competition information.
// This entity is shared between Consensus and Blockchain.
type Competition interface {
	// Gets the current competition branch up to the block number
	// specified. If a previous head has been invalidated, then
	// switchHeads will be true.
	Branch(headBlockNumber uint64) (branch []Block, rootID int, switchHeads bool)
}