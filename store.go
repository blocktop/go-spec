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

type Store interface {
	OpenBlock(blockNumber uint64) (StoreBlock, error)
	StoreBlock() StoreBlock 
	Close()
	GetRoot() string
	Put(context.Context, Marshalled) error
	Get(ctx context.Context, hash string, obj Marshalled) error
	TreeGet(ctx context.Context, key string, obj Marshalled) error
	Hash(data []byte, links Links) (string, error)
}

type StoreBlock interface {
	IsOpen() (bool, uint64)
	Submit(context.Context, Block) (root string, err error)
	GetRoot() string
	Commit(context.Context) error
	Revert() error
	TreePut(ctx context.Context, key string, obj Marshalled) error
	TreeGet(ctx context.Context, key string, obj Marshalled) error
}
