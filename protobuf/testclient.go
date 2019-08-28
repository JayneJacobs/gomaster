package protobuf

import (
	"flag"
	"gotrain/goMaster/protobuf/commlayer"
	"gotrain/goMaster/protobuf/protout"
	"log"
	"strings"
)

// Ex gives the address of a server and port and creats the flags
func Ex() {
	op := flag.String("type", "", "Server (s) or client ()c ?")
	address := flag.String("addr", ":8080", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("Received: ", msg)
	}
}

func runClient(dest string) {
	c := commlayer.NewConnection(commlayer.Protobuf)
	ship := &protout.Ship{
		Shipname:    "Jayne's Dream",
		CaptainName: "Jala",
		Crew: []*protout.Ship_CrewMember{
			&protout.Ship_CrewMember{1, "Jim", 5, "bass"},
			&protout.Ship_CrewMember{2, "Jayne", 4, "Engineer"},
			&protout.Ship_CrewMember{3, "Scott", 6, "Keyboard"},
		},
	}
	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error While Sending message", err)
	}
}
