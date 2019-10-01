package bandchat

import (
	"fmt"
	"gotrain/goMaster/blogger"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//Logger uses hlogger

// Run takes a connection string and returns an error to start a server
func Run(connection string) error {
	logger := blogger.GetInstance("serverRun")
	l, err := net.Listen("tcp", connection)
	logger.Println("Just in Run function")

	if err != nil {
		logger.Println("Error connecting to chat client", err)
		fmt.Println("Error connecting to chat client", err)
		//return err
	}
	logger.Println("Just in Run function past error")
	r := CreateRoom("BandChat")

	go func(l net.Listener) {
		logger := blogger.GetInstance("go routine Listener")
		logger.Println("Just in Run go routine")
		for {
			logger.Println("J go routine for loop")
			conn, err := l.Accept()
			if err != nil {
				logger.Println("Error accepting connection from chat client", err)
				//break
			}
			logger.Println("Just before handle")
			go handleConnection(r, conn)
		}
	}(l)

	go func() {
		// Handle SIGINT and SIGTERM.
		logger := blogger.GetInstance("goRoutein2")
		logger.Println("Just befor sigterm")
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()
	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error accepting connection from chat client", err)
			break
		}
		go handleConnection(r, conn)
	}
	return err
}

func handleConnection(r *room, c net.Conn) {
	logger := blogger.GetInstance("serverHandle")
	logger.Println("Received request from client", c.RemoteAddr())
	r.AddClient(c)
}
