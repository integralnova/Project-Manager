package models

import (
	"log"
)

func (m *Datatings) NewPermit(permit PermitModelPermitID) error {
	stmt := `INSERT INTO permitid (permitID)
	VALUES (?)`
	_, err := m.DB.Exec(stmt, permit.Permit)
	log.Println(err)
	return err
}

func (m *Datatings) UpdatePermitCompany(permit PermitModelPermitCompany) error {
	stmt := `INSERT INTO permit_company (permit, companyName)
	VALUES (?, ?)`
	_, err := m.DB.Exec(stmt, permit.Permit, permit.CompanyName)
	log.Println(err)
	return err
}

func (m *Datatings) UpdatePermitDesigner(permit PermitModelPermitDesigner) error {
	stmt := `INSERT INTO permit_designer (permit, designer, dateStarted, dateFinished)
	VALUES (?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, permit.Permit, permit.Designer, permit.DateStarted, permit.DateCompleted)
	log.Println(err)
	return err
}

func (m *Datatings) UpdatePermitDateReceived(permit PermitModelPermitDateReceived) error {
	stmt := `UPSERT permit_designer SET dateFinished = ? WHERE permit = ?`
	_, err := m.DB.Exec(stmt, permit.DateReceived, permit.Permit)
	log.Println(err)
	return err
}

func (m *Datatings) Getpermits() ([]PermitModelPermitID, error) {
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
