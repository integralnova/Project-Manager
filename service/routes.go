package service

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	//mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /permits/create", app.newPermit)
	mux.HandleFunc("GET /", app.getpermits)
	mux.HandleFunc("GET /permit/{id}", app.getpermit)
	return mux

}
