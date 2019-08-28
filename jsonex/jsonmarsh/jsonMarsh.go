package jsonmarsh

import (
	"encoding/json"
	"fmt"
)

// BandMember lists Name Age PhoneNumer and email
type BandMember struct {
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age,omitempty"`
	PhoneNumber string   `json:"phoneNumber,omitempty"`
	Email       string   `json:"email,omitempty"`
	AccessCodes []string `json:"accessCodes,omitempty"`
}

// Instrument gives the name Key Octaves and list of players
type Instrument struct {
	Name    string
	Key     string
	Octaves int
	Players []BandMember
}

// MarshalEx Structs
func MarshalEx() {

	umj := BandMember{"Jim", 64, "215-519-0733", "jrpadilla@comcast.net", []string{"ONO", "AAA"}}
	umjjson, err := json.Marshal(&umj)
	if err != nil {
		fmt.Println(err)
	}

	// Marshal embedded structs as a slice of Players
	fmt.Println(string(umjjson))
	lmj := BandMember{"Jim", 64, "215-519-0733", "jrpadilla@comcast.net", []string{}}
	lmjjson, err := json.Marshal(&lmj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(lmjjson))

	fmj := BandMember{"Jayne", 61, "856-275-3638", "jaynejacobs@comcast.net", []string{}}
	fmjjson, err := json.Marshal(&fmj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(fmjjson))
	ifmj := Instrument{"Guitar", "A", 3, []BandMember{lmj, fmj}}
	ifmjjson, err := json.Marshal(&ifmj)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ifmjjson))
	// Marshal maps
	mi := map[string]int{"Jim": 3, "Jayne": 5}
	il, err := json.Marshal(&mi)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(il))

}
