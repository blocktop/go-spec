package spec

import (
	"github.com/golang/protobuf/proto"
)

type NetworkNode interface {
	GetPeerID() string
	Broadcast(proto.Message, *MessageProtocol)
	OnMessageReceived(MessageReceivedHandler)
	Sign(data []byte) ([]byte, error)
	Verify(peerID string, data []byte, sig []byte) (bool, error)
}

type MessageReceivedHandler func(proto.Message, *MessageProtocol)
