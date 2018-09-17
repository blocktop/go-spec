package spec

import "context"

type Blockchain interface {
	GetType() string
	ProduceGenesisBlock()
	GetBlockGenerator() BlockGenerator
	GetConsensus() Consensus
	Start(context.Context)
	Stop()
	IsRunning() bool
	GetReceiveChan() chan<- *BroadcastBlock
	GetBroadcastChan() <-chan *BroadcastBlock
	GetConfirmChan() <-chan Block
	GetConfirmLocalChan() <-chan Block
}
