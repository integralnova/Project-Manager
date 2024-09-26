package service

import (
	"html/template"
	"net/http"
	"time"

	"github.com/integralnova/Project-Manager/models"
	"github.com/integralnova/Project-Manager/permit_tracker/sqlite"
)

func (app *app) getpermit(w http.ResponseWriter, r *http.Request) {

}

func (app *app) getpermits(w http.ResponseWriter, r *http.Request) {
	permits, err := app.permits.Getpermits()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles("./assets/templates/permits.html")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	p := make([]sqlite.PermitsViewModel, len(permits))
	for i, _ := range permits {
		p[i] = sqlite.TranslatePermit(permits[i])
	}

	t.Execute(w, map[string]any{"permits": p})
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

	dr, _ := time.Parse("2006-01-02", r.PostForm.Get("date_received"))
	dd, _ := time.Parse("2006-01-02", r.PostForm.Get("date_due"))

	permit := models.PermitsModel{
		PermitID:     r.PostForm.Get("permit_id"),
		CompanyName:  r.PostForm.Get("company_name"),
		Reference:    r.PostForm.Get("reference"),
		DateReceived: dr,
		DateDue:      dd,
		PermitStatus: r.PostForm.Get("permit_status"),
		Designer:     r.PostForm.Get("name"),
	}

	err = app.permits.Insert(permit)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	http.Redirect(w, r, "/permits", http.StatusFound)
}
