package bandchat

import (
	"fmt"
	"gotrain/goMaster/blogger"
	"net"
	"sync"
)

type room struct {
	name    string
	Msgch   chan string
	clients map[chan<- string]struct{}
	Quit    chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room {
	r := &room{
		name:    name,
		Msgch:   make(chan string),
		RWMutex: new(sync.RWMutex),
		clients: make(map[chan<- string]struct{}),
		Quit:    make(chan struct{}),
	}
	r.Run()
	return r
}

func (r *room) AddClient(c net.Conn) {
	var Logger = blogger.GetInstance("roomAddClient")
	Logger.Println("Adding client", c.RemoteAddr())
	r.Lock()
	wc, done := StartClient(r.Msgch, c, r.Quit)
	r.clients[wc] = struct{}{}
	r.Unlock()

	//remove client when done is signalled
	go func() {
		<-done
		r.RemoveClient(wc)
	}()
}

func (r *room) ClCount() int {
	return len(r.clients)
}

func (r *room) RemoveClient(wc chan<- string) {
	var Logger = blogger.GetInstance("roomsRemovingclient")
	Logger.Println("Removing client ", r.name)
	r.Lock()
	close(wc)
	delete(r.clients, wc)
	r.Unlock()
	select {
	case <-r.Quit:
		if len(r.clients) == 0 {
			close(r.Msgch)
		}
	default:
	}
}

func (r *room) Run() {
	var Logger = blogger.GetInstance("roomRun")
	Logger.Println("Starting chat room", r.name)

	Logger.Println("In go routine", r.name)
	for msg := range r.Msgch {
		r.broadcastMsg(msg)
		Logger.Printf("Starting chat room line67 %v \n", msg)
	}

	Logger.Println("After go routine", r.name)

}

func (r *room) broadcastMsg(msg string) {
	var Logger = blogger.GetInstance("broadcastclient")
	Logger.Println("broadcastMsge Line81 in room", r.name)
	r.RLock()
	defer r.RUnlock()
	fmt.Println("Received message: ", r.name)
	for wc, _ := range r.clients {
		go func(wc chan<- string) {
			wc <- msg
		}(wc)
	}
}
