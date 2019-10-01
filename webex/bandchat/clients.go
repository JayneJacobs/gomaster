package bandchat

import (
	"bufio"
	"gotrain/goMaster/blogger"
	"net"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

func StartClient(msgCh chan<- string, cn net.Conn, quit chan struct{}) (chan<- string, <-chan struct{}) {
	logger := blogger.GetInstance("clientsStartClient")
	c := new(client)
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)
	done := make(chan struct{})

	//setup the reader
	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			msgCh <- scanner.Text()
		}
		done <- struct{}{}
	}()

	//setup the writer
	c.writeMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()
		case <-done:
		}
	}()

	return c.wc, done
}

func (c *client) writeMonitor() {
	go func() {
		logger := blogger.GetInstance("clientswriteMonitor")
		for s := range c.wc {
			logger.Println("Sending", s)
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}