package spec

import (
	"github.com/golang/protobuf/proto"
)

type Network interface {
	Broadcast(proto.Message, *MessageProtocol)
	OnMessageReceived(MessageReceivedHandler)
}

type MessageReceivedHandler func(proto.Message, *MessageProtocol)
