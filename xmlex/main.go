package main

import (
	"gotrain/goMaster/jsonex/jsonmarsh"
	"gotrain/goMaster/xmlex/xmldecode"
	"gotrain/goMaster/xmlex/xmlencode"
	"gotrain/goMaster/xmlex/xmlmarsh"
	"gotrain/goMaster/xmlex/xmlunmarsh"
)

func main() {
	xmlmarsh.MarshalEx()
	xmlencode.EncodeEx()
	xmlunmarsh.Ex()
	xmldecode.Ex(jsonmarsh.BandMember{}, "file.xml")
}
