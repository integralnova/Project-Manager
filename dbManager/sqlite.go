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
	fmt.Println(sql.Drivers())
	db, err := sql.Open("sqlite3", "./permits.db")
	checkErr(err)
	defer db.Close()
	fmt.Println(sqlited.Version())

	sqlStmt := `
	CREATE TABLE permits ( id INTEGER PRIMARY KEY, permit INTEGER);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	sql2 := `
	INSERT INTO permits VALUES 1;
	`
	x, err := db.Exec(sql2)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	fmt.Println(x)

}
