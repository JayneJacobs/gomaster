package dblayer

import (
	"database/sql"
	// _ mysql is use fro a driver
	_ "github.com/go-sql-driver/mysql"
)

// MySQLDataStore ...
type MySQLDataStore struct {
	*sql.DB
}

// NewMySQLDataStore returns a NewMySQL connection
func NewMySQLDataStore(conn string) (*MySQLDataStore, error) {
	db, err := sql.Open("mysql", conn)
	if err != nil {
		return nil, err
	}
	return &MySQLDataStore{
		DB: db}, nil
}

// AddMember ...
func (jsql *MySQLDataStore) AddMember(bm *BandMember) error {
	_, err := jsql.Exec("INSERT INTO Personnel (Name,Instrument,Position) VALUES(?,?,?)", bm.Name, bm.Instrument, bm.Position)
	return err
}

// FindMember ...
func (jsql *MySQLDataStore) FindMember(id int) (BandMember, error) {
	row := jsql.QueryRow("Select * from Personnel wehre id = ?", id)
	bm := BandMember{}
	err := row.Scan(&bm.ID, &bm.Name, &bm.Instrument, &bm.Position)
	return bm, err
}

// AllMembers ...
func (jsql *MySQLDataStore) AllMembers() (Band, error) {
	rows, err := jsql.Query("Select* from Personnel;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	members := Band{}
	for rows.Next() {
		member := BandMember{}
		err = rows.Scan(&member.ID, &member.Name, &member.Instrument, &member.Position)
		if err != nil {
			members = append(members, member)
		}
	}
	err = rows.Err()
	return members, err
}
