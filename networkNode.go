package spec

import (
	"context"

	"github.com/golang/protobuf/proto"
)

type NetworkNode interface {
	Bootstrap(context.Context) error
	GetPeerID() string
	Listen(context.Context)
	Broadcast(ctx context.Context, message proto.Message, from string, protocol *MessageProtocol)
	Sign(data []byte) ([]byte, error)
	Verify(peerID string, data []byte, sig []byte, pubKey []byte) (bool, error)
	GetReceiveChan() <-chan *NetworkMessage
	Close()
}
