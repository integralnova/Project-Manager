package service

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/integralnova/Project-Manager/models"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	permits *models.Datatings
}

func WebService() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	app := app{
		permits: &models.Datatings{
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
