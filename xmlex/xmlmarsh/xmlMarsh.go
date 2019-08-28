package xmlmarsh

import (
	"encoding/xml"
	"fmt"
)

// MakeBandMember allows the use of the BandMember xml struct Name, Age, PhoneNumber
type MakeBandMember interface {
	MarshalBand(b *BandMember)
}

// BandMember lists Name Age PhoneNumer and email
type BandMember struct {
	Name        string   `xml:"name,omitempty,attr"`
	Age         int      `xml:"age,omitempty"`
	PhoneNumber string   `xml:"phoneNumber,omitempty"`
	Email       string   `xml:"email,omitempty"`
	AccessCodes []string `xml:"accessCodes>code,omitempty"` //specify sublist
}

// Instrument gives the name Key Octaves and list of players
type Instrument struct {
	XMLName xml.Name     `xml:"instrument"`
	Name    string       `xml:"name,attr"`
	Key     string       `xml:"key"`
	Octaves int          `xml:"octaves"`
	Players []BandMember `xml:"players>member"`
}

// MarshalEx Structs
func MarshalEx() {
	umj := BandMember{"Rich", 64, "215 - 555 - 1212", "rheim@comcast.net", []string{"GUT", "SNG"}}
	smj := BandMember{"Elvis", 62, "215 - 555 - 1212", "", []string{"DRM", "SNG"}} //empty field ommited

	ifmj := Instrument{Name: "Guitar", Key: "G", Octaves: 3, Players: []BandMember{umj, smj}}
	fmjxml, err := xml.Marshal(&ifmj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fmjxml))
	fmjxmlIN, err := xml.MarshalIndent(&ifmj, " ", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fmjxmlIN))

	// Marshal slice
	s := []int{1, 2, 3, 4}
	fmjxml, err = xml.MarshalIndent(&s, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(xml.Header, string(fmjxml))

	// Marshal Maps
	// mi := map[string]int{"Jim": 3, "Jayne": 5}
	// il, err := xml.Marshal(&mi)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(il))

}

// MarshalIndent takes an input filename and output file name. It will read
// the contents of the file and beautify it.
func MarshalIndent(filein string, fileout string) {

}
