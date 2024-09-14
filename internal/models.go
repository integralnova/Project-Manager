package models

import "time"

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
