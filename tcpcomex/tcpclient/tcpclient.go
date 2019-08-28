package tcpclient

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// RunClient takes an address string and returns an error
func RunClient(address string) error {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return err
	}

	defer conn.Close()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("What message would you like to send? ")

	for scanner.Scan() && err == nil {
		fmt.Println("Writing ", scanner.Text())
		conn.Write(append(scanner.Bytes(), '\r'))

		fmt.Println("What message would you like to send?")
		buffer := make([]byte, 1024)
		//conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		_, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(buffer))
	}
	return err
}
