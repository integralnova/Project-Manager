package models

import "time"

type PermitsModel struct {
	ID           int
	PermitID     string
	CompanyName  string
	Reference    string
	DateReceived time.Time
	DateDue      time.Time
	PermitStatus permitstatus
	Owner        employee
}

type permitstatus string

const (
	StatusPending   permitstatus = "Pending"
	StatusApproved  permitstatus = "Approved"
	StatusRejected  permitstatus = "Rejected"
	StatusCancelled permitstatus = "Cancelled"
)

type designstatus string

const (
	DesignPending      designstatus = "Design Pending"
	PoleForemanPending designstatus = "Pole Foreman pending"
	WOPending          designstatus = "WO pending"
	QCPending          designstatus = "QC pending"
)

type employee struct {
	id        int
	firstname string
}
