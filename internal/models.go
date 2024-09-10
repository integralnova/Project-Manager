package models

import "time"

type Post struct {
	id               int
	permitid         string
	companyName      string
	reference        string
	datereceived     time.Time
	dateapproved     string
	designer         employee
	designstartdate  time.Time
	designfinishdate time.Time
	permittype       string
	permitstatus     permitstatus
	owner            employee
	status           designstatus
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