package models

import "time"

//deprecate
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
