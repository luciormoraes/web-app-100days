package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// TODO
func ConnectDB() *sql.DB {
	// conn:="user dbname password host sslmode"
	conn := "user=root dbname=test_db password=root host=0.0.0.0 sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
