package commlayer

import (
	"gotrain/goMaster/protobuf/protout"
)

const (
	// Protobuf is a uint8
	Protobuf uint8 = iota
)

// HydraConnnection is an inteace that takes an empty interface and a destination
// string and returns an error.
type HydraConnnection interface {
	EncodeAndSend(obj interface{}, destination string) error
	ListenAndDecode(listenaddress string) (chan interface{}, error)
}

// NewConnection allows access to NewProtoHanler
func NewConnection(connType uint8) HydraConnnection {
	switch connType {
	case Protobuf:
		return protout.NewProtoHandler()
	}
	return nil
}
