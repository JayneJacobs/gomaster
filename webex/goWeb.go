package main

import (
	"crypto/md5"
	"flag"
	"gotrain/goMaster/blogger"
	"gotrain/goMaster/webex/bandchat"
	"gotrain/goMaster/webex/bandportal"
	"gotrain/goMaster/webex/dbex/dblayer/vault"
	"strings"
)

func main() {

	//code to populate some password in the pass vault
	db, err := vault.ConnectPasswordVault()
	if err != nil {
		return
	}
	minapss := md5.Sum([]byte("minaspass"))
	jaynepass := []byte("brainstorm") //Not hashed
	jimpass := md5.Sum([]byte("jimspass"))
	caropass := md5.Sum([]byte("carospass"))
	vault.AddBytesToVault(db, "Mina", minapss[:])
	vault.AddBytesToVault(db, "Jim", jimpass[:])
	vault.AddBytesToVault(db, "Caro", caropass[:])
	vault.AddBytesToVault(db, "jayne@comcast.net", jaynepass[:])
	db.Close()

	///////RunChat Server below
	logger := blogger.GetInstance("WebServers")
	logger.Println("Starting Web Service")
	operation := flag.String("o", "w", "Operation: w for web \n c for chat")
	flag.Parse()
	switch strings.ToLower(*operation) {
	case "c":
		err := bandchat.Run(":2100")
		if err != nil {
			logger.Println("Could not run band web portal", err)
		}
	case "w":
		err := bandportal.Run()
		if err != nil {
			logger.Println("Could not run hydra web portal", err)
		}
	}

}
