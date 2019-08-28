package protout

import (
	"io/ioutil"
	"time"

	"errors"
	"log"
	"net"

	"github.com/golang/protobuf/proto"
)

//ProtoHandler is an empty struct
type ProtoHandler struct{}

// NewProtoHandler is a Constructor that returns a pointer to Protohandler type
func NewProtoHandler() *ProtoHandler {
	return new(ProtoHandler)
}

// EncodeAndSend sends data to ProtoHandler type from the Encode
func (psend *ProtoHandler) EncodeAndSend(obj interface{}, dest string) error {
	v, ok := obj.(*Ship) //Convert Empty Interface to Concrete *Ship Concrete type
	if !ok {
		return errors.New("Proto: Unknown message type")
	}
	data, err := proto.Marshal(v)
	if err != nil {
		return err
	}
	return sendmessage(data, dest)
}

// ListenAndDecode takes an address string and returns a channel sends  data to *ProtoHandler type .
func (psend *ProtoHandler) ListenAndDecode(listenaddress string) (chan interface{}, error) {
	outChan := make(chan interface{})
	l, err := net.Listen("tcp", listenaddress)
	if err != nil {
		return outChan, err
	}
	log.Println("Listening to ", listenaddress)
	go func() {
		defer l.Close()

		for {
			c, err := l.Accept()
			if err != nil {
				break
			}
			log.Println("Accepted Connection from ", c.RemoteAddr())
			go func(c net.Conn) {
				defer c.Close()
				for {
					buffer, err := ioutil.ReadAll(c)
					if err != nil {
						break // exit the for loop and the go routine
					}
					if len(buffer) == 0 {
						continue
					}
					obj, err := psend.DecodeProto(buffer)
					if err != nil {
						continue
					}
					select {
					case outChan <- obj: // Pass obj to outChan
					case <-time.After(1 * time.Second):
					default:
					}
				}
			}(c) //Conn object type from l.Accept()
		}

	}()
	return outChan, nil
}

// EncodeProto takes an interface an dreturns a byte string and error
func EncodeProto(obj interface{}) ([]byte, error) {
	if v, ok := obj.(*Ship); ok {
		return proto.Marshal(v)
	}
	err := errors.New("Proto: Unknown message type")
	return nil, err
}

// DecodeProto takes a byte array and returns a byte string and error
func (psend *ProtoHandler) DecodeProto(buffer []byte) (*Ship, error) {
	pb := new(Ship)
	err := proto.Unmarshal(buffer, pb)
	return pb, err
}

func sendmessage(buffer []byte, dest string) error {
	conn, err := net.Dial("tcp", dest)
	if err != nil {
		return err
	}
	defer conn.Close()
	log.Printf("Sending %d bytes to %s \n", len(buffer), dest)
	_, err = conn.Write(buffer)
	return err
}
