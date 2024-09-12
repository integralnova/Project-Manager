package dbmanager

import (
	"database/sql"
	"log"

	models "github.com/integralnova/Project-Manager/internal"
)

type PermitModel struct {
	DB *sql.DB
}

func (m *PermitModel) Insert(permit models.PermitsModel) error {
	stmt := `INSERT INTO posts (permitID, companyName, reference, dateReceived, dateDue, permitStatus, designer)
	VALUES (?, ?, ?, ?, ?, ?, ?, datetime('now'))`
	_, err := m.DB.Exec(stmt, permit.PermitID, permit.CompanyName, permit.Reference, permit.DateReceived, permit.DateDue, permit.PermitStatus, permit.Designer)
	log.Println(err)
	return err
}

// this could get ugly. Need to figure out if i can some how break all permits in to pages
// Look into limit offset
func (m *PermitModel) Getpermits() ([]models.PermitsModel, error) {
	stmt := `SELECT * FROM permits ORDER BY id DESC`
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
