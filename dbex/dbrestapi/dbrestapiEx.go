package dbrestapi

import (
	"gotrain/goMaster/Hydra/hydraconfigurator"
	"log"
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
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "../../webex/apiconfig.json")
	if err != nil {
		return err
	}
	h := NewBandReqHandler()
	err = h.Connect(conf.DB, conf.Conn)
	if err != nil {
		log.Fatal("Error connecting to db", err)
	}
	http.HandleFunc("/band/", h.HandleBandRequests)
	// http.Handle("/band/", h)//If uncomenting SERVEHTTP version
	return nil
}

// RunAPI Initlializes server
func RunAPI() {
	InitializeAPIHandlers()
	http.ListenAndServe(":8061", nil)
}
