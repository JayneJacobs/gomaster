package csvencoder

import (
	"encoding/csv"
	"fmt"
	"gotrain/goMaster/csvex/csvmarsh"
	"os"
)

// Ex ...
func Ex() {
	f, err := os.Create("file.csv")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	cm := csvmarsh.BandMember{Name: "Scott", Age: 60, Email: "scooper@gmail.com", AccessCodes: []string{"Keys", "Drums"}}
	lbm := []csvmarsh.BandMember{cm}
	si := csvmarsh.Instrument{Name: "Sax", Key: "G", Octaves: 3, Players: lbm}
	err = csv.NewEncoder(f).Encode(&si)
	if err != nil {
		fmt.Println(err)
	}

}
