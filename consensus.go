package spec

type Consensus interface {
	AddBlock(block Block) (added bool)
	GetBestBranch() []Block
	WasSeen(block Block) bool
	SetCompeted(head Block)
	OnBlockConfirmed(BlockConfirmedHandler)
}

type BlockComparator func([]Block) Block
type BlockConfirmedHandler func(Block)
