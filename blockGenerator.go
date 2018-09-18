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

type BlockGenerator interface {
	GetType() string
	Unmarshal(proto.Message, map[string]TransactionHandler) Block
	LogTransaction(txn Transaction)
	UnlogTransactions(txns []Transaction)
	GenerateGenesisBlock() Block
	GenerateBlock(parentBranch []Block) (newBlock Block)
}
