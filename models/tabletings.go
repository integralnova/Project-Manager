package models

import (
	"log"
)

func (m *Datatings) UpdatePermit(permit PermitModelPermitID) error {
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

func (m *Datatings) GetPermitById(id int) (PermitModelPermitID, error) {
	stmt := `SELECT * FROM permitid WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	p := PermitModelPermitID{}
	err := row.Scan(&p.ID, &p.Permit)
	if err != nil {
		return PermitModelPermitID{}, err
	}

	return p, nil
}

func (m *Datatings) GetPermitByName(name string) (PermitsModel, error) {
	stmt := `SELECT * FROM permit_company WHERE companyName = ?`
	row := m.DB.QueryRow(stmt, name)

	p := PermitsModel{}
	err := row.Scan(&p.PermitID, &p.CompanyName, &p.Reference, &p.DateReceived, &p.DateDue, &p.PermitStatus, &p.Designer)
	if err != nil {
		return PermitsModel{}, err
	}

	return p, nil
}

/* TODO

GetPermitId
GetPermitById
GetPermitCompany

UpdatePermitCompany
UpdatePermitDesigner
UpdatePermitDateReceived
updatePermitDateDue
updatePermitDateSubmitted

*/
