package service

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/integralnova/Project-Manager/models"
)

var templates = []string{
	"./assets/templates/permits.html",
}

func (app *app) getpermit(w http.ResponseWriter, r *http.Request) {

}

func (app *app) getpermits(w http.ResponseWriter, r *http.Request) {
	permits, err := app.permits.Getpermits()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t.Execute(w, map[string]any{"permits": permits})
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
	fmt.Println("new permit")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	dr, _ := time.Parse("2006-01-02", r.PostForm.Get("date_received"))
	dd, _ := time.Parse("2006-01-02", r.PostForm.Get("date_due"))

	permit := models.NewPermitsModel()
	permit.PermitID = isblank(r.PostForm, "permit_id", permit.PermitID)
	permit.CompanyName = isblank(r.PostForm, "company_name", permit.CompanyName)
	permit.Reference = isblank(r.PostForm, "reference", permit.Reference)
	permit.DateReceived = dr
	permit.DateDue = dd
	permit.PermitStatus = isblank(r.PostForm, "permit_status", permit.PermitStatus)
	permit.Designer = isblank(r.PostForm, "name", permit.Designer)

	err = app.permits.InsertPermit(permit)

	if err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/permits", http.StatusFound)
}
func isblank(form map[string][]string, key, defaultValue string) string {
	if value := form[key]; len(value) > 0 && len(value[0]) > 0 {
		return value[0]
	}
	return defaultValue
}
