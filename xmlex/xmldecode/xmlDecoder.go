package xmldecode

import (
	"encoding/xml"
	"fmt"
	"os"
)

//Ex  decode xml from a file
func Ex(i interface{}, filename string) {

	err := decodeXML(&i, filename)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
	}
}

//DecodeJSON takes interface and filename and returns an error
// It decodes the contents of a file to an empty struct
func decodeXML(i *interface{}, filename string) error {
	fmt.Println("Decoding json")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	return xml.NewDecoder(file).Decode(&i)

}
