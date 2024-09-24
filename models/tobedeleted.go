package models

import "log"

func (m *Datatings) GetPermitByCompanyName(name string) (PermitsModel, error) {
	stmt := `SELECT * FROM permit_company WHERE companyName = ?`
	row := m.DB.QueryRow(stmt, name)

	p := PermitsModel{}
	err := row.Scan(&p.PermitID, &p.CompanyName, &p.Reference, &p.DateReceived, &p.DateDue, &p.PermitStatus, &p.Designer)
	if err != nil {
		return PermitsModel{}, err
	}

	return p, nil
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