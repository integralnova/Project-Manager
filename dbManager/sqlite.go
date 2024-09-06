package dbmanager

import (
	"database/sql"
	"fmt"
	"log"

	sqlited "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Connect to database
func makedb() {
	db, err := sql.Open("sqlite3", "./permits.db")
	checkErr(err)
	defer db.Close()
	fmt.Println(sqlited.Version())
}
