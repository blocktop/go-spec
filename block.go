package blckit_interface

// Block interface specifies getters required for `blckit` to function.
// Custom blocks must implement this interface.
type Block interface {
	// returns the identifier of the block
	GetID() string

	// returns the identifier of the block's parent
	GetParentID() string

	// returns the sequential number of the block in the blockchain
	GetBlockNumber() uint64
}
