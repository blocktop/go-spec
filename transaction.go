package spec

import "github.com/golang/protobuf/proto"

type Transaction interface {
	GetType() string
	GetVersion() string
	GetID() string
	Marshal() proto.Message
	Unmarshal(proto.Message)
}
