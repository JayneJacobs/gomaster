package xmlencode

import (
	"encoding/xml"
	"fmt"
	"gotrain/goMaster/xmlex/xmlmarsh"
	"os"
)

// EncodeEx ...
func EncodeEx() {
	f, err := os.Create("file.xml")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	cm := xmlmarsh.BandMember{Name: "Scott", Age: 60, Email: "scooper@gmail.com", AccessCodes: []string{"Keys", "Drums"}}
	lbm := []xmlmarsh.BandMember{cm}
	si := xmlmarsh.Instrument{Name: "Sax", Key: "G", Octaves: 3, Players: lbm}

	enc := xml.NewEncoder(f)
	enc.Indent(" ", "  ")
	err = enc.Encode(si)
	if err != nil {
		fmt.Println(err)
	}

}
