package spec

type Blockchain interface {
	GetType() string
	ProduceGenesisBlock()
	GetBlockGenerator() BlockGenerator
	GetConsensus() Consensus
	Start()
	Stop()
	IsRunning() bool
	ReceiveBlock(block Block) 
	OnBlockGenerated(BlockGeneratedHandler)
	OnBlockConfirmed(BlockConfirmedHandler)
	OnLocalBlockConfirmed(BlockConfirmedHandler)
}

type BlockGeneratedHandler func(Block)
