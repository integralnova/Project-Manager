package models

import (
	"time"

	"database/sql"
)

type Datatings struct {
	DB *sql.DB
}

type PermitsModel struct {
	ID           int
	PermitID     string
	CompanyName  string
	Reference    string
	DateReceived time.Time
	DateDue      time.Time
	PermitStatus string
	Designer     string
}

type PermitsViewModel struct {
	ID           int
	PermitID     string
	CompanyName  string
	Reference    string
	DateReceived string
	DateDue      string
	PermitStatus string
	Designer     string
}

func TranslatePermit(permit PermitsModel) PermitsViewModel {
	return PermitsViewModel{
		ID:           permit.ID,
		PermitID:     permit.PermitID,
		CompanyName:  permit.CompanyName,
		Reference:    permit.Reference,
		DateReceived: permit.DateReceived.Format("2006-01-02"), // Format the date
		DateDue:      permit.DateDue.Format("2006-01-02"),      // Format the date
		PermitStatus: permit.PermitStatus,
		Designer:     permit.Designer,
	}
}

func NewPermitsModel() PermitsModel {
	return PermitsModel{
		PermitID:     "empty",
		CompanyName:  "empty",
		Reference:    "empty",
		DateReceived: time.Now(),
		DateDue:      time.Now().AddDate(0, 1, 60),
		PermitStatus: "empty",
		Designer:     "empty",
	}
}
