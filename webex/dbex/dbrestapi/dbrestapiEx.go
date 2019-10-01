package dbrestapi

import (
	"fmt"
	"gotrain/goMaster/webex/configurate"
	"net/http"
)

// DBlayerconfig ...
type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

// InitializeAPIHandlers ...
func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	fmt.Println("In Initialize")
	err := configurate.GetConfiguration(configurate.JSON, conf, "./apiconfig.json")
	if err != nil {
		fmt.Println("just past GetConfiguration", err)
	}
	h := NewBandReqHandler()
	fmt.Println("Past NewReqHandler")
	err = h.Connect(conf.DB, conf.Conn)
	if err != nil {
		err = fmt.Errorf("Error connecting to db %s", err)
	}
	http.HandleFunc("/band/", h.HandleBandRequests)
	fmt.Println("Past HandleFunc")
	// http.Handle("/band/", h)//If uncomenting SERVEHTTP version
	return nil
}

// RunAPI Initlializes server
func RunAPI() error {
	InitializeAPIHandlers()
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("This is a RunAPI error", err)
		return err
	}
	return nil
}
