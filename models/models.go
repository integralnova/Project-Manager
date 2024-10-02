package models

import (
	"time"

	"database/sql"
)

type Datatings struct {
	DB *sql.DB
}

// deprecate
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

type PermitModelPermitID struct {
	ID     int
	Permit string
}
type PermitModelPermitCompany struct {
	Permit      string
	CompanyName string
}

type PermitModelPermitDesigner struct {
	Permit        string
	Designer      string
	DateStarted   time.Time
	DateCompleted time.Time
}

type PermitModelPermitDateReceived struct {
	Permit       string
	DateReceived time.Time
}

type PermitModelPermitDateDue struct {
	Permit  string
	DateDue time.Time
}

type PermitModelPermitDateSubmit struct {
	Permit        string
	DateSubmitted time.Time
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