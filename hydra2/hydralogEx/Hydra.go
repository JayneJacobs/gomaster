package hydra

import (
	"fmt"
	"net/http"

	"gotrain/goMaster/hydra/hlogger"
)

// Hydra will take a string and log start of the server
func Hydra(s string, l string) {
	logger := hlogger.GetInstance("Server")
	logger.Printf("This is the URL for %v:  %v", s, l)

	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance("sroot")
	homeMessage := "Welcome to the API Test Tool"
	fmt.Fprint(w, homeMessage)

	logger.Println("Received an API Get Request")
}
