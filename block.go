// Copyright © 2018 J. Strobus White.
// This file is part of the blocktop blockchain development kit.
//
// Blocktop is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Blocktop is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with blocktop. If not, see <http://www.gnu.org/licenses/>.

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
