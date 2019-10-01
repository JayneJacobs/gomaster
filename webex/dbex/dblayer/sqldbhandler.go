package dblayer

import (
	"database/sql"
	"fmt"

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
	_, err := jsql.Exec("INSERT INTO Personnel (Name,SecurityClearance,Position) VALUES(?,?,?)", bm.Name, bm.SecurityClearance, bm.Position)
	return err
}

// FindMember ...
func (jsql *MySQLDataStore) FindMember(id int) (BandMember, error) {
	row := jsql.QueryRow("Select * from Personnel where id = ?", id)
	bm := BandMember{}
	err := row.Scan(&bm.ID, &bm.Name, &bm.SecurityClearance, &bm.Position)
	return bm, err
}

// AllMembers ...
func (jsql *MySQLDataStore) AllMembers() (Band, error) {
	rows, err := jsql.Query("Select* from Personnel;")
	fmt.Println("In All members function")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	members := Band{}
	for rows.Next() {
		member := BandMember{}
		err = rows.Scan(&member.ID, &member.Name, &member.SecurityClearance, &member.Position)

		if err == nil {
			members = append(members, member)
		}
		fmt.Println(members)
	}

	err = rows.Err()
	fmt.Println(err, members)
	return members, err
}
