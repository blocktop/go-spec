package spec

// Competition represents the latest block competition information.
// This entity is shared between Consensus and Blockchain.
type Competition interface {
	// Branches returns the current competing branches being tracked
	// in the consensus system indexed by RootID.
	Branches() map[int]CompetingBranch
}

type CompetingBranch interface {
	// Blocks returns the blocks contained in the branch, ordered from
	// head to confirming root.
	Blocks() []Block

	// RootID is the identifier of the branch.
	RootID() int

	// ConsecutiveLocalHits returns the number of times the branch has
	// received blocks for evaluation from only the local node. This
	// is an indicator that the current node is operating in an echo
	// chamber.
	ConsecutiveLocalHits() int

	// HitRate is a measure of popularity of the branch. If returns the
	// number of blocks per second that are being evaluated from the
	// network and locally for addition to the branch.
	HitRate() float64
}
