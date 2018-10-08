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

// Block interface specifies getters required for `blocktop` to function.
// Custom blocks must implement this interface.
type Block interface {
	Marshalled 

	// returns the identifier of the block's parent
	ParentHash() string

	// returns the sequential number of the block in the blockchain
	BlockNumber() uint64

	// validates the block
	Valid() bool

	// gets the block's transactions
	Transactions() []Transaction

	// gets the Unix time in milliseconds the block was generated
	Timestamp() int64
}
