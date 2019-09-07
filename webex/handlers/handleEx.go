package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
)

type testhandler struct {
	r int
	n string
}

// NewHandler provides handlers
func NewHandler() testhandler {
	return testhandler{
		r: rand.Int(),
		n: "Jayne",
	}
}

// ServeHttp satisfies the interface requirements of server.go inerits the handler interface
func (h testhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "Welcome to my handler type example")
	case "/testhandle":
		fmt.Fprintf(w, "test handle object wit random number %v", h.r)
	case "/myname":
		fmt.Fprintf(w, "Hello %v", h.n)
	}
	fmt.Println(r.URL.Query())

}
