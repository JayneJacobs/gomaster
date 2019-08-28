package main

import (
	"gotrain/goMaster/jsonex/jsondecoder"
	"gotrain/goMaster/jsonex/jsonencoder"
	"gotrain/goMaster/jsonex/jsonmarsh"
	"gotrain/goMaster/jsonex/jsonunmarsh"
)

func main() {
	jsonmarsh.MarshalEx()
	jsonencoder.Ex()
	jsonunmarsh.Ex()
	jsondecoder.Ex(jsonmarsh.BandMember{}, "file.json")
}
