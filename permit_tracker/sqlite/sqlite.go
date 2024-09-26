package sqlite

import (
	"database/sql"
	"log"

	models "github.com/integralnova/Project-Manager/models"
)

type PermitModel struct {
	DB *sql.DB
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

func TranslatePermit(permit models.PermitsModel) PermitsViewModel {
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



func (m *PermitModel) Insert(permit models.PermitsModel) error {

	stmt := `INSERT INTO permits (permitID, companyName, reference, dateReceived, dateDue, permitStatus, designer)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, permit.PermitID, permit.CompanyName, permit.Reference, permit.DateReceived.Format("2006-01-02"), permit.DateDue.Format("2006-01-02"), permit.PermitStatus, permit.Designer)

	log.Println(err)
	return err
}

// time.Time parases incorrectly from datetime in sqlite
func (m *PermitModel) Getpermits() ([]models.PermitsModel, error) {
	stmt := `SELECT * FROM permits ORDER BY id ASC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	permits := []models.PermitsModel{}

	for rows.Next() {
		p := models.PermitsModel{}
		err := rows.Scan(&p.ID, &p.PermitID, &p.CompanyName, &p.Reference, &p.DateReceived, &p.DateDue, &p.PermitStatus, &p.Designer)
		if err != nil {
			return nil, err
		}
		permits = append(permits, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return permits, nil
}

func (m *PermitModel) CreatePermit(title string) error {
	stmt := `INSERT INTO permits (title)
	VALUES (?)`
	_, err := m.DB.Exec(stmt, title)
	return err
}

func (m *PermitModel) ChangeDateReceived(date string) error {
	stmt := `UPDATE permits SET dateReceived = ?`
	_, err := m.DB.Exec(stmt, date)
	return err
}
