package service

import (
	"html/template"
	"net/http"
	"time"

	models "github.com/integralnova/Project-Manager/internal"
)

func (app *app) getpermits(w http.ResponseWriter, r *http.Request) {
	posts, err := app.permits.Getpermits()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./templates/permits.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, map[string]any{"Posts": posts})
}

func (app *app) createPost(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/post.create.page.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, nil)
}

func (app *app) newPermit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	permit := models.PermitsModel{
		PermitID:     r.PostForm.Get("permit_id"),
		CompanyName:  r.PostForm.Get("company_name"),
		Reference:    r.PostForm.Get("reference"),
		DateReceived: time.Now(),
		DateDue:      time.Now().AddDate(0, 0, 30),
		PermitStatus: r.PostForm.Get("permit_status"),
		Designer:     r.PostForm.Get("name"),
	}

	err = app.permits.Insert(permit)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
