package hydracommlayer

import (
	"gotrain/goMaster/Hydra/hydracomms/hydramessages/hydraproto"
)

const (
	// Protobuf is a uint8
	Protobuf uint8 = iota
)

type HydraConnnection interface {
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

func NewConnection(connType uint8) HydraConnection {
	switch connType {
	case Protobuf:
		return hydraproto.NewProtoHandler()
	}
	return nil
}
