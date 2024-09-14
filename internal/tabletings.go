package models

import (
	"database/sql"
	"fmt"
	"log"
)

type PermitModel struct {
	DB *sql.DB
}

func (m *PermitModel) NewPermit(permit PermitModelPermitID) error {
	stmt := `INSERT INTO permitid (permitID)
	VALUES (?)`
	_, err := m.DB.Exec(stmt, permit.Permit)
	log.Println(err)
	return err
}

func (m *PermitModel) UpdatePermitCompany(permit PermitModelPermitCompany) error {
	stmt := `INSERT INTO permit_company (permit, companyName)
	VALUES (?, ?)`
	_, err := m.DB.Exec(stmt, permit.Permit, permit.CompanyName)
	log.Println(err)
	return err
}

func (m *PermitModel) UpdatePermitDesigner(permit PermitModelPermitDesigner) error {
	stmt := `INSERT INTO permit_designer (permit, designer, dateStarted, dateFinished)
	VALUES (?, ?, ?, ?)`
	a, err := m.DB.Exec(stmt, permit.Permit, permit.Designer, permit.DateStarted, permit.DateCompleted)
	log.Println(err)
	a1, _ := a.LastInsertId()
	a2, _ := a.RowsAffected()
	fmt.Print(a1, a2)
	return err
}

func (m *PemritModel) UpdatePermit

func (m *PermitModel) Getpermits() ([]PermitModelPermitID, error) {
	stmt := `SELECT * FROM permitid ORDER BY id ASC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	permits := []PermitModelPermitID{}

	for rows.Next() {
		p := PermitModelPermitID{}
		err := rows.Scan(&p.ID, &p.Permit)
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
