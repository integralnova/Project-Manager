package models

import "time"

type Permits struct {
	ID               int
	PermitID         string
	CompanyName      string
	Reference        string
	DateReceived     time.Time
	DateApproved     string
	Designer         employee
	DesignStartDate  time.Time
	DesignFinishDate time.Time
	PermitType       string
	PermitStatus     permitstatus
	Owner            employee
	Status           designstatus
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
