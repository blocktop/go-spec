package spec

import "context"

type Consensus interface {
	Start(context.Context)
	AddBlock(block Block, local bool) (added bool)
	WasSeen(block Block) bool
	SetCompeted(head Block)
	GetConfirmChan() <-chan Block
	GetCompetitionChan() <-chan []Block
}

type BlockComparator func([]Block) Block
