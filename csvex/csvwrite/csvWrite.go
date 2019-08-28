package csvwrite

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// BandMember lists Name Age PhoneNumer and email
type BandMember struct {
	Name        string   `csv:"name,omitempty"`
	Age         int      `csv:"age,omitempty"`
	PhoneNumber string   `csv:"phoneNumber,omitempty"`
	Email       string   `csv:"email,omitempty"`
	AccessCodes []string `csv:"accessCodes,omitempty"`
}

// Instrument gives the name Key Octaves and list of players
type Instrument struct {
	Name    string
	Key     string
	Octaves int
	Players []BandMember
}

// Exwrite a csvfile
func Exwrite() {
	records := [][]string{
		{"Jaro", "5", "ALA,IOI"},
		{"Jayne", "5", "ALA,IOI"},
		{"Rich", "5", "ALA,IOI"},
	}

	file, err := os.Create("cfilenew.csv")

	if err != nil {

		fmt.Println(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(records)
	w.Comma = ';'
	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatal(err)
	// 	}

	w.Flush()
	err = w.Error()
	if err != nil {
		log.Fatal(err)
	}
}

// Ex takes an input json filepath and delivers a .csv file
func Ex(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	// record := csv.NewReader(file)
	// records, err := record.Read()
	if err != nil {
		fmt.Println(err)
	}
	outfile, err := os.Create("outfile.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	w := csv.NewWriter(outfile)
	w.Comma = ';'

	// for _, line := range records {
	// 	if err := w.WriteAll(records); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("Writing this line", line)
	// }

}
