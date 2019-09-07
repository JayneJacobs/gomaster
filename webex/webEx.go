package webex

import (
	"fmt"
	"gotrain/goMaster/hlogger"
	"gotrain/goMaster/webex/handlers"
	"net/http"
)

func Run() {
	http.HandleFunc("/", sroot)
	http.Handle("/testhandle", handlers.NewHandler()) //Uses ServeHTTP method
	http.Handle("/myname", handlers.NewHandler())     //Uses ServeHTTP method

	http.HandleFunc("/testquery", queryTestHandler)
	//http.ListenAndServe(":8080", handlers.NewHandler()) //change to default mux for query
	http.ListenAndServe(":8080", nil) //change to default mux for query
	//
	// server := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      handlers.NewHandler(),
	// 	ReadTimeout:  5 * time.Second,
	// 	WriteTimeout: 5 * time.Second,
	// }
	// server.ListenAndServe()

}

func queryTestHandler(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	// log.Println("Forms", r.Form)
	q := r.URL.Query()
	message := fmt.Sprintf("Query map: %v \n", q)

	v1, v2 := q.Get("key1"), q.Get("key2")
	name := q.Get("myname")
	if name != "nil" {
		fmt.Fprintf(w, "Hello %s", message)
	}
	if v1 == v2 {
		message = message + fmt.Sprintf("v1 an dV2 are equal %s \n", v1)
	} else {
		message = message + fmt.Sprint("V1 is %s, V2 is %s \n", v1, v2)
	}
	fmt.Fprint(w, message)

}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance(string("URL"))
	fmt.Fprint(w, "welcome to the Hydra software sytem", r.Body, r.URL)
	logger.Println(r.Body)

}
