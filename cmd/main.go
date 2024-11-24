package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {

	connStr := "user=postgres password=467912 dbname=web_storage sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
