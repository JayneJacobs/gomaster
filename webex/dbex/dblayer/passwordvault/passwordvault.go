package passwordvault

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

var errNilDB = errors.New("Database handler is nil....")

func ConnectPasswordVault() (*bolt.DB, error) {
	fmt.Println("passwordvault connect")
	db, err := bolt.Open("band.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("passwordvault connected")
	return db, nil
}

func AddToVault(db *bolt.DB, username, password string) error {
	if db == nil {
		return errNilDB
	}
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PasswordVault"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(username), []byte(password))
		return err
	})
}

func AddBytesToVault(db *bolt.DB, username string, password []byte) error {
	if db == nil {
		return errNilDB
	}
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PasswordVault"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(username), password)
		return err
	})
}

func GetPassword(db *bolt.DB, username string) (string, error) {

	fmt.Println("passwordvault 1")
	if db == nil {
		return "", errNilDB
	}
	password := ""
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PasswordVault"))
		v := b.Get([]byte(username))
		password = string(v)
		return nil
	})
	fmt.Println("passwordvault 2")
	return password, err
}

func GetPasswordBytes(db *bolt.DB, username string) ([]byte, error) {
	fmt.Println("passwordvault 3")
	if db == nil {
		return nil, errNilDB
	}
	fmt.Println("passwordvault 4")
	password := []byte{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PasswordVault"))
		password = b.Get([]byte(username))
		return nil
	})

	return password, err
}
