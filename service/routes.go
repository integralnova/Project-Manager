package service

import "net/http"

func (app *app) routes() http.Handler {
	mux := http.NewServeMux()
	//mux.HandleFunc("GET /", app.getHome)
	mux.HandleFunc("GET /posts/create", app.createPost)
	mux.HandleFunc("POST /permit/create", app.newPermit)
	mux.HandleFunc("GET /permits", app.getpermits)
	return mux
}
