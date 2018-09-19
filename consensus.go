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

import "context"

type Consensus interface {
	Start(ctx context.Context, confirmChan chan<- Block, confirmLocalChan chan<- Block, competeChan chan<- []Block)
	AddBlock(block Block, local bool) (added bool)
	WasSeen(block Block) bool
	SetCompeted(head Block)
}

type BlockComparator func([]Block) Block
