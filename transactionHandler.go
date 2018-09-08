package spec

import "github.com/golang/protobuf/proto"

type TransactionHandler interface {
	GetType() string
	Unmarshal(proto.Message) Transaction
	Execute(Transaction) (ok bool)
}
