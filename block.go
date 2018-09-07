package spec

import "github.com/golang/protobuf/proto"

// Block interface specifies getters required for `blckit` to function.
// Custom blocks must implement this interface.
type Block interface {
	// returns the type of the block
	GetType() string

	// returns the version of the block
	GetVersion() string

	// returns the identifier of the block
	GetID() string

	// returns the identifier of the block's parent
	GetParentID() string

	// returns the sequential number of the block in the blockchain
	GetBlockNumber() uint64

	// validates the block
	Validate() bool

	// gets the block's transactions
	GetTransactions() []Transaction

	// gets the Unix time in milliseconds the block was generated
	GetTimestamp() int64

	// converts the block to a protocolbuf
	Marshal() proto.Message

	Unmarshal(proto.Message, map[string]TransactionHandler)
}
