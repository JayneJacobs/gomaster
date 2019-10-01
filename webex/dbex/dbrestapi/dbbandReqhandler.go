package dbrestapi

import (
	"encoding/json"
	"fmt"
	"gotrain/goMaster/dbex/dblayer"
	"net/http"
	"strconv"
)

// BandReqHandler loads the db struct
type BandReqHandler struct {
	dbConn dblayer.DBLayer
}

// NewBandReqHandler returns a Band Request Handler for the DB
func NewBandReqHandler() *BandReqHandler {
	fmt.Println("returning BandReqHandler")
	return new(BandReqHandler)

}

// Connect ...
func (bdwreq *BandReqHandler) Connect(o, conn string) error {
	dblayer, err := dblayer.ConnectDatabase(o, conn)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println("Connected to db")
	bdwreq.dbConn = dblayer
	return nil
}

/*
func (bdwreq BandReqHandler) ServeHttp(w http.ResponseWriter, r *http.Request) {
	switch r.Method(
	case "GET":
		ids := r.RequestURI[len("/band/"):]
		id, err := strconv(Atoi(ids))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "id %s provided is not of valid number. \n", ids)
			return
		}
		bm, err :- bdwreq.dbConn.FindMamber(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "id %s occured when search for id %d . \n", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&bm)
	case "POST":
		bm := new(dblayer.BandMember)
		err := json.NewDecoder(r.Body).Decode(bm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error %s reported. \n", ids)
            return
		}
		bdwreq.dbConn.AddMember(bm)
	)
}
*/

// HandleBandRequests ...
func (bdwreq *BandReqHandler) HandleBandRequests(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In HandleBandRequests")
	switch r.Method {
	case "GET":
		fmt.Println("in case Get")
		ids := r.RequestURI[len("/findmember/"):]
		id, err := strconv.Atoi(ids)
		fmt.Printf("This is the id. %v ", id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Printf("This is the id in error. %v ", id)
			fmt.Fprintf(w, "id %s provided is not of valid number. \n", ids)
			return
		}
		bm, err := bdwreq.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "id %s occured when search for id %d . \n", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&bm)
	case "POST":
		bm := new(dblayer.BandMember)
		err := json.NewDecoder(r.Body).Decode(bm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error %s reported. \n", err)
			return
		}
		err = bdwreq.dbConn.AddMember(bm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// 	fmt.Fprintf(w, "Error %s occured whike adding member to database", err)
			return
		}
		fmt.Fprintf(w, "Successfully added id %d \n", bm.ID)
	default:
		fmt.Println("No correct case")
	}

}