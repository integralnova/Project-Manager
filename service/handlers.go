package service

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
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

	//whatever it works
	if dr.Year() <= 2010 {
		dr = time.Now()
	}
	if dd.Year() <= 2010 {
		dd = dr.AddDate(0, 0, 90)
	}

	form := r.PostForm
	permit := models.NewPermitsModel()
	permit.PermitID = isblank(form, "permit_id", permit.PermitID).(string)
	permit.CompanyName = isblank(form, "company_name", permit.CompanyName).(string)
	permit.Reference = isblank(form, "reference", permit.Reference).(string)
	permit.DateReceived = dr
	permit.DateDue = dd
	permit.PermitStatus = isblank(form, "permit_status", permit.PermitStatus).(string)
	permit.Designer = isblank(form, "name", permit.Designer).(string)

	err = app.permits.InsertPermit(permit)

	if err != nil {
		fmt.Println(err)
		return
	}

	http.Redirect(w, r, "/permits", http.StatusFound)
}
func isblank(form url.Values, key string, defaultVal any) any {
	if value := form[key]; len(value) > 0 && len(value[0]) > 0 {
		return value[0]
	}
	return defaultVal
}
