package models

import (
	"database/sql"
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
