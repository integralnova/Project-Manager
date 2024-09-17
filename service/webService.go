package service

import (
	"database/sql"
	"log"
	"net/http"

	sqlite "github.com/integralnova/Project-Manager/permit_tracker/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	permits *sqlite.PermitModel
}

func WebService() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		permits: &sqlite.PermitModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listening on :8000")
	srv.ListenAndServe()
}
