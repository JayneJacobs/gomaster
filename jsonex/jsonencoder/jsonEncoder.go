package jsonencoder

import (
	"encoding/json"
	"fmt"
	"gotrain/goMaster/jsonex/jsonmarsh"
	"os"
)

// Ex ...
func Ex() {
	f, err := os.Create("file.json")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	cm := jsonmarsh.BandMember{Name: "Scott", Age: 60, Email: "scooper@gmail.com", AccessCodes: []string{"Keys", "Drums"}}
	lbm := []jsonmarsh.BandMember{cm}
	si := jsonmarsh.Instrument{Name: "Sax", Key: "G", Octaves: 3, Players: lbm}
	err = json.NewEncoder(f).Encode(&si)
	if err != nil {
		fmt.Println(err)
	}

}
