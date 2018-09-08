package spec

import "github.com/golang/protobuf/proto"

type BlockGenerator interface {
	GetType() string
	Unmarshal(proto.Message, map[string]TransactionHandler) Block
	LogTransaction(txn Transaction)
	UnlogTransactions(txns []Transaction)
	ProduceGenesisBlock() Block
	GenerateBlock(parentBranch []Block) (newBlock Block)
}
