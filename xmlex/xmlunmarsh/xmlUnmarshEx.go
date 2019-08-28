package xmlunmarsh

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"gotrain/goMaster/jsonex/jsonmarsh"
	"gotrain/goMaster/xmlex/xmlmarsh"
	"io/ioutil"
	"os"
)

// Ex ...
func Ex() {
	f, err := ioutil.ReadFile("file.xml")
	if err != nil {
		fmt.Println(err)
	}
	// sbyte := []byte(f)
	si := new(xmlmarsh.Instrument)
	xml.Unmarshal((f), si) //unmarshal document into struct
	fmt.Println(si.Key, si.Octaves, si.Players)

	m := make(map[int]string)

	xml.Unmarshal(f, &m)

	mapfile, err := os.OpenFile("mapdata.xml", os.O_CREATE|os.O_APPEND, 0444)
	defer mapfile.Close()

	if err != nil {
		fmt.Println(err)
	}
	mi := jsonmarsh.Instrument{}

	err = xml.Unmarshal(f, &mi)
	if err != nil {
		fmt.Println(err)
	}
	wb := bufio.NewWriter(mapfile)
	err = xml.Unmarshal(f, &mi)
	success, err := wb.WriteString("This will be the struct")
	if err != nil {
		fmt.Println(err, success)
	}
	fmt.Println(success)
}
