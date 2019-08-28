package tcpserver

import (
	"bufio"
	"flag"
	"fmt"
	"gotrain/goMaster/tcpcomex/tcpclient"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Ex takes one address string and a flag and runs either a client or server
func Ex() {
	op := flag.String("type", "value", "Server (s) or client(c)?")
	address := flag.String("addr", ":8000", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		tcpclient.RunClient(*address)
	}
}

func runServer(address string) error {
	// for ship := range hydramessages.ListaneAndDecode(hydramessages.GOB, address)
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()
	log.Println("Listening")
	for {
		c, err := l.Accept()
		if err != nil {
			return err
		}
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	writer := bufio.NewWriter(c)

	for {
		//buffer := make([]byte, 1024)
		c.SetReadDeadline(time.Now().Add(5 * time.Second))
		line, err := reader.ReadString('\r') //bufio.Reader allows ReadString method by passing c
		//+, err := c.Read(buffer)
		if err != nil && err != io.EOF {
			log.Println(err)
			return
		} else if err == io.EOF {
			log.Println("Connection closed")
			return
		}
		fmt.Printf("Received %s from address %s \n", line[:len(line)-1], c.RemoteAddr())
		writer.WriteString("Message received... ")
		writer.Flush()
	}
}
