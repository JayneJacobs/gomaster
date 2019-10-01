package dblayer

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDataStore ...
type MongoDataStore struct {
	*mgo.Session
}

// NewMongoStore sets up the Mongo db connection
func NewMongoStore(conn string) (*MongoDataStore, error) {
	log.Println(conn)
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}
	return &MongoDataStore{Session: session}, nil
}

// AddMember ...
func (ms *MongoDataStore) AddMember(bm *BandMember) error {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	// bm = BandMember{}
	// err := membership.Find(bson.M{"id": id}).One(&bm)
	return personnel.Insert(bm)

}

// FindMember ...
func (ms *MongoDataStore) FindMember(id int) (BandMember, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	bm := BandMember{}
	err := personnel.Find(bson.M{"id": id}).One(&bm)
	return bm, err
}

// AllMembers ...
func (ms *MongoDataStore) AllMembers() (Band, error) {
	session := ms.Copy()
	defer session.Close()
	personnel := session.DB("Hydra").C("Personnel")
	members := Band{}
	err := personnel.Find(nil).All(&members)
	return members, err
}
