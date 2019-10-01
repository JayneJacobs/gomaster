package main

import (
	"crypto/md5"
	"gotrain/goMaster/webex/dbex/dblayer/vault"
)

func main() {
	db, err := vault.ConnectPasswordVault()

	if err != nil {
		return
	}
	jaynepss := md5.Sum([]byte("brainstorm"))
	minapss1 := md5.Sum([]byte("minaspass1"))
	minapss2 := md5.Sum([]byte("minaspass2"))
	vault.AddBytesToVault(db, "jayne@comcast.net", jaynepss[:])
	vault.AddBytesToVault(db, "Mina1", minapss1[:])
	vault.AddBytesToVault(db, "Mina2", minapss2[:])
	db.Close()
}
