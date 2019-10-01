package dblayer

import (
	"errors"
	"log"
)

const (
	//MONGO ...
	MONGO = "mongodb"
	// MYSQL ...
	MYSQL = "mysql"
)

var errtype = errors.New("That is not the right Database")

// DBLayer provides a struct for any Database
type DBLayer interface {
	AddMember(bm *BandMember) error
	FindMember(id int) (BandMember, error)
	AllMembers() (Band, error)
}

// Band ...
type Band []BandMember

// BandMember Defines a Band Member position and instrument
type BandMember struct {
	ID                int    `json:"id"  bson:"id"`
	Name              string `json:"name" bson:"name"`
	SecurityClearance int    `json:"securityclearance" bson:"securityclearance"`
	Position          string `json:"position" bson:"position"`
}

// ConnectDatabase returns the type of database depending on what is used
func ConnectDatabase(o string, cstring string) (DBLayer, error) {
	switch o {
	case MONGO:
		return NewMongoStore(cstring)
	case MYSQL:
		return NewMySQLDataStore(cstring)
	}
	log.Panicln("Could not find that Database type ", o)
	return nil, errtype
}
