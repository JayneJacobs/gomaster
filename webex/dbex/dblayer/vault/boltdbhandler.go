package vault

import (
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

var (
	// ErrNilDB Database error
	ErrNilDB = errors.New("Error Nil DB")
)

// ConnectPasswordVault takes a pointer to a DB struct
func ConnectPasswordVault() (*bolt.DB, error) {
	fmt.Println("InConnect")
	db, err := bolt.Open("band.db", 0600, nil)
	if err != nil {
		fmt.Println("error in connect to Vault")
		return nil, err

	}

	fmt.Println("Sending back the connection to DB.")
	return db, nil
}

//AddBytesToVault takes a pointer to a db Connection, username and password string and returns an error
func AddBytesToVault(db *bolt.DB, username string, password []byte) error {
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PasswordVault"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(username), password) // Enter a key and a value
		return err
	})
}

// AddToVault takes a pointer to a database and returns and error
func AddToVault(db *bolt.DB, username, password string) error {
	if db == nil {
		return ErrNilDB
	}
	return db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("PasswordVault"))
		if err != nil {
			return err
		}
		err = b.Put([]byte(username), []byte(password)) // Enter a key and a value
		return err
	})
}

// GetPassword takes a pointer to a DB connection and a username string.
// Returns a password and and error
func GetPassword(db *bolt.DB, username string) (string, error) {
	fmt.Println("I n GetPasswords")
	if db == nil {
		return "This is what is the Error", ErrNilDB
	}
	password := ""
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PasswordVault"))
		v := b.Get([]byte(username))
		password = string(v)
		fmt.Printf("This is returned from the db %s, as a %T \n", password, password)
		return nil
	})
	fmt.Printf("This is the error %v", password)
	return password, err
}

// GetPasswordBytes takes a pointer to a DB connection and a username string.
// Returns a password []bytes and and error
func GetPasswordBytes(db *bolt.DB, username string) ([]byte, error) {
	fmt.Println("In GetPasswordsBytes")
	if db == nil {
		fmt.Println("No DB")
		return nil, ErrNilDB

	}
	password := []byte{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("PasswordVault"))
		if b == nil {
			fmt.Println("No PasswordVault bucket found")
			return errors.New("No PasswordVault bucket found")
		}
		password = b.Get([]byte(username))
		return nil

	})

	return password, err
}
