package blckit_interface

type ConsensusSpec interface {
	ShouldRemoveBlock(b Block, maxBlockNumberSeen uint64) bool
}

