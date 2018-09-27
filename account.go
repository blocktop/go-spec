package spec

import "github.com/golang/protobuf/proto"

type Account interface {
	Address() string
	Marshal() proto.Message
}
