package hydrachat

import (
	"gotrain/goMaster/hydra/hlogger"

	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var logger = hlogger.GetInstance()

//Run returns an error and listens on port 2100
func Run(connection string) error {
	l, err := net.Listen("tcp", connection)
	r := CreateRoom("HydraChat")
	logger.Printf("Past createroom")
	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}
	go func(l net.Listener) {
		logger.Printf("netListener go routine")
		for {
			conn, err := l.Accept()
			if err != nil {
				logger.Println("Error accepting connection from chat client", err)
				//break
			}
			go handleConnection(r, conn)
		}
	}(l)

	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		logger.Printf("Past second go routine")
		//l.Close()
		fmt.Println("Closing tcp connection")
		//close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		//os.Exit(0)
	}()
	return nil
}

func handleConnection(r *room, c net.Conn) {
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
