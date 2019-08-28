package jsondecoder

import (
	"encoding/json"
	"fmt"
	"os"
)

//Ex  decode json from a file
func Ex(i interface{}, filename string) {

	err := decodeJSON(&i, filename)
	fmt.Println(i)
	if err != nil {
		fmt.Println(err)
	}
}

//DecodeJSON takes interface and filename and returns an error
// It decodes the contents of a file to an empty struct
func decodeJSON(i *interface{}, filename string) error {
	fmt.Println("Decoding json")
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	return json.NewDecoder(file).Decode(&i)

}
