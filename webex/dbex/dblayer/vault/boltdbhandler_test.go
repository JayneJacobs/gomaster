package vault

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

// TestGetPassword ...
func TestGetPassword(t *testing.T) {
	db, err := ConnectPasswordVault()
	defer db.Close()
	//put below into a test
	// op := flag.String("op", "", "Add or Get operation?") // A pointer
	// user := flag.String("u", "", "Enter username")
	// pass := flag.String("p", "", "Enter password")
	// flag.Parse()

	op := "GET"
	user := "jayne@comcast.net"
	fmt.Println("Enter the hostname of the vde_cm: ")
	var pass string

	fmt.Scan(&pass)
	switch strings.ToUpper(op) { //dereference the value
	case "Add":
		if len(user) != 0 && len(pass) != 0 {
			err = AddToVault(db, user, pass)
			if err != nil {
				log.Fatal(err)
			}
		}
	case "GET":
		if len(user) != 0 {
			pass, err := GetPassword(db, user)
			if err != nil {
				log.Fatal(err)
			}

			t.Logf("Found password: %v \n ", pass)
		}
	}
}
