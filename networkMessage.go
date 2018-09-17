package spec

import "github.com/golang/protobuf/proto"

type NetworkMessage struct {
	Message  proto.Message
	Protocol *MessageProtocol
	From		 string
}
