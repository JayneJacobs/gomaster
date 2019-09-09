package restapi

import (
	"gotrain/goMaster/Hydra/hydraconfigurator"
	"gotrain/goMaster/dbex/dbrestapi"
	"log"
	"net/http"
)

// DBlayerconfig  is the DB struct
type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

// InitializeAPIHandlers Provides a new handler
func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "../../apiconfig.json")
	if err != nil {
		log.Println("Error decoding Json", err)
		return err
	}
	h := dbrestapi.NewBandReqHandler() //Consructor for the BandReqHandler s
	err = h.Connect(conf.DB, conf.Conn)
	if err != nil {
		log.Panicln("Error connecting to db", err)
		return err
	}
	http.HandleFunc("/band/", h.HandleBandRequests)
	// http.Handle("/band/", h)//If uncomenting SERVEHTTP version
	return nil
}

// RunAPI Initlializes server
func RunAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}

	return http.ListenAndServe(":8061", nil)

}
