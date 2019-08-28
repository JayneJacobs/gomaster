package jsonunmarsh

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gotrain/goMaster/jsonex/jsonmarsh"
	"io/ioutil"
	"os"
)

// Ex ...
func Ex() {
	f, err := ioutil.ReadFile("file.json")
	if err != nil {
		fmt.Println(err)
	}
	// sbyte := []byte(f)
	si := new(jsonmarsh.Instrument)
	json.Unmarshal((f), si)
	fmt.Println(si.Key, si.Octaves, si.Players)
	m := make(map[int]string)
	data := []byte(`{"1":"Jim", "2":"Jayne"}`)
	json.Unmarshal(data, &m)
	mapfile, err := os.OpenFile("mapdata.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer mapfile.Close()
	wb := bufio.NewWriter(mapfile)

	wb.WriteString(fmt.Sprintln("This is the  map data ", string(data)))

}
