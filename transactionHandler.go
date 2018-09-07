package spec

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

type TransactionHandler interface {
	GetType() string
	UnmarshalAny(*any.Any) Transaction
	Unmarshal(proto.Message) Transaction
	Execute(Transaction) (ok bool)
}
