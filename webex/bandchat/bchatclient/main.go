package main

import (
	"bufio"
	"fmt"
	"gotrain/goMaster/blogger"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	//"band/bandConfigurator"
)

//Logger uses blogger

func main() {
	var Logger = blogger.GetInstance("clientmain")
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))
	fmt.Println("Starting bandChatClient....")
	Logger.Println("Starting bandChatClient....")
	fmt.Println("What's your name?")
	fmt.Scanln(&name)

	/*confStruct := struct{
		Name string `name:"name"`
		RemoteAddr string `name:"remoteip"`
		TCP bool `name:"tcp"`
	}{}

	bandConfigurator.GetConfiguration(bandConfigurator.CUSTOM,&confStruct,"chat.conf")
	name = confStruct.Name
		proto := "tcp"
	if !confStruct.TCP{
		proto = "udp"
	}
	*/

	fmt.Printf("Hello %s, connecting to the band chat system.... \n", name)
	Logger.Printf("Hello %v, connecting to the band chat system.... \n", name)
	conn, err := net.Dial("tcp", "localhost:2300")
	if err != nil {
		log.Fatal("Could not connect to band chat system", err)
	}
	fmt.Println("Connected to band chat system")
	Logger.Println("Connected to band chat system")
	name += ":"
	defer conn.Close()
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			Logger.Println("InBufioScanner loop")
			Logger.Println(scanner.Text())
		}
	}()

	/*
		for err==nil {
			msg := ""
			fmt.Print(name)
			fmt.Scan(&msg)
			msg = name+msg+"\n"
			fmt.Println("Duplicate: " + msg)
			_,err = fmt.Fprintf(conn,msg)

		}
	*/

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()                     //Get string representation of a buffer
		_, err = fmt.Fprintf(conn, name+msg+"\n") //Writes message to tcp.
		Logger.Printf("This is the message%v%v\n", name, msg)
	}
}
