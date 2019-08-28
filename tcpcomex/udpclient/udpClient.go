package udpclient

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

// RunUDPClient takes an address string and returns an error it starts teh client
func RunUDPClient(address string) error {
	conn, err := net.Dial("udp", address)
	if err != nil {
		return nil
	}
	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)

	filebytes, err := ioutil.ReadFile("inputfile.csv")

	if err != nil {
		log.Fatal(err)
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(filebytes)))
	base64.StdEncoding.Encode(dst, filebytes)
	_, err = conn.Write(dst)

	log.Println("Sending ", len(dst), " bytes")
	fmt.Println("What message would you like to send?")

	for scanner.Scan() && err == nil {

		fmt.Println("Writing ", scanner.Text())
		conn.Write(scanner.Bytes()) //bytes converts the input text to a slice of bytes

		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer)) //convert buffer to a string
		fmt.Println("What message would you like to send?")
	}
	return err
}
