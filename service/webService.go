package service

import (
	"database/sql"
	"log"
	"net/http"

	sqlite "github.com/integralnova/Project-Manager/dbManager"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	posts *sqlite.PostModel
}

func WebService() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		posts: &sqlite.PostModel{
			DB: db,
		},
	}

	srv := http.Server{
		Addr:    ":8000",
		Handler: app.routes(),
	}

	log.Println("Listing on :8000")
	srv.ListenAndServe()
}
