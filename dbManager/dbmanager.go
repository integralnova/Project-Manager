package dbmanager

import (
	"fmt"
	"time"
)

type employee struct {
	id        int
	firstname string
}

var Employees = make(map[int]employee)

func NewEmployee(employee employee) {
	Employees[employee.id] = employee
}

func GetEmployee() map[int]employee {
	return Employees
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

type Permit struct {
	id               int
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

func DbManager() {
	// dbManager logic
}

type permits map[int]Permit

func init() {
	// dbManager init logic
	p := make(permits)
	GetPermits(p)
	makedb()
}

func sqlite() {
	// sqlite logic
}

func GetPermits(p permits) {
	// GetPermits logic
	for _, permit := range p {
		fmt.Println(permit)
	}
}

func NewPermit(permit Permit) {
	// NewPermit logic
}
