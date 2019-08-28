package udpserver

import (
	"encoding/base64"
	"flag"
	"fmt"
	"gotrain/goMaster/tcpcomex/udpclient"
	"log"
	"net"
	"os"
	"strings"
)

// Ex gives access to the server/client functions this provides a chat program after encoding a file
func Ex() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8000", "address? host:port")
	flag.Parse()
	switch strings.ToUpper(*op) {
	case "S":
		runUDPServer(*address)
	case "C":
		udpclient.RunUDPClient(*address)
	}
}

func runUDPServer(address string) error {
	pc, err := net.ListenPacket("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	buffer64 := make([]byte, 4096)

	fmt.Println("Listenting...")
	n, _, err := pc.ReadFrom(buffer64) //received data
	if err != nil {
		log.Fatal(err)
	}
	dst64 := make([]byte, base64.StdEncoding.DecodedLen(n)) //buffer for received data
	_, err = base64.StdEncoding.Decode(dst64, buffer64[:n])
	if err != nil {
		log.Fatal(err)
	}
	file, _ := os.Create("outputfile.csv")
	file.Write(dst64)
	err = file.Close()
	for {
		buffer := make([]byte, 1024)
		_, addr, _ := pc.ReadFrom(buffer)
		fmt.Printf("Received %s from address %s \n", string(buffer), addr)
		_, err := pc.WriteTo([]byte("Message Received"), addr)
		if err != nil {
			log.Fatal("Could not write back on connection ", err)
			return file.Close()
		}
	}

}
