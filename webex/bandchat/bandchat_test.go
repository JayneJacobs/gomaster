package bandchat

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func() {
		t.Log("Startign Bandchat Server")
		if err := Run(":2300"); err != nil {
			t.Error("Could not start chat server ", err)
			return
		}
		t.Log("Started chat server")
	}
}
func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode...")

	}
	t.Log("Testing band chat sernd in short mode....")
	go once.Do(chatServerFunc(t))

	time.Sleep(1 * time.Second)
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("Anonymous%d", rand.Intn(400))

	t.Logf("Hello %s, connecting to the chat syste... \n", name)
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		t.Fatal("Could not connect to chat system because port is wrong. Test Passed", err)
	}
	t.Log("connected to the chat system")
	name += ":" //name:my chat message
	defer conn.Close()
	msgCh := make(chan string)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			recvmsg := scanner.Text()
			sentmsg := <-msgCh
			if strings.Compare(recvmsg, sentmsg) != 0 {
				t.Errorf("Chat Message %s does not match %s", recvmsg, sentmsg)
			}
		}
	}()

	for i := 0; i < 10; i++ {
		msgbody := fmt.Sprintf("RandomMessage %d", rand.Intn(400))
		msg := name + msgbody
		_, err = fmt.Fprint(conn, msg+"\n")
		// Anonymous4:Randommessage1
		if err != nil {
			t.Error(err)
			return
		}
		msgCh <- msg
	}

}

func TestServerConnection(t *testing.T) {
	t.Log("Test band chat receive messages....")
	f := chatServerFunc(t)
	go once.Do(f)
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "localhost:2300")
	if err != nil {
		t.Fatal("Could not connect to chat server", err)
	}
	conn.Close()
}
