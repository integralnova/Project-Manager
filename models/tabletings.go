package models

import (
	"log"
)

// How to deal with duplicates.
// Find at creation. use UPSERT
// Might have duplicate permit names
// Insert permit
func (m *Datatings) InsertPermit(permit PermitModelPermitID) error {
	stmt := `INSERT INTO permitid (permitID)
	VALUES (?)`
	_, err := m.DB.Exec(stmt, permit.Permit)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Insert permit_company
func (m *Datatings) InsertPermitCompany(permit PermitModelPermitCompany) error {
	stmt := `INSERT INTO permit_company (permit, companyName)
	VALUES (?, ?)`
	_, err := m.DB.Exec(stmt, permit.Permit, permit.CompanyName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// All permits
// NEEDS rewrite
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

// find permit by id WHY?
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

// find company by permit
func (m *Datatings) GetCompanyByPermit(company string) (PermitModelPermitCompany, error) {
	stmt := ""
	row := m.DB.QueryRow(stmt, company)

	p := PermitModelPermitCompany{}
	err := row.Scan(&p.Permit, &p.CompanyName)
	if err != nil {
		return PermitModelPermitCompany{}, err
	}

	return p, nil
}

// Some fields in Permit Designer need to be nil sometimes. Or need new table and model. Design finish date not always available//
// https://stackoverflow.com/questions/43854117/detecting-uninitialized-struct-fields
// Does inserting with default fields in the struct mess with db?
func (m *Datatings) InsertDesigner(p PermitModelPermitDesigner) error {
	stmt := `INSERT INTO permit_designer (permit, designer, dateStarted, dateCompleted) VALUE (?,?,?,?)`
	_, err := m.DB.Exec(stmt, p.Permit, p.Designer, p.DateStarted, p.DateCompleted)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// get designer. Uh what about when date completed is null
func (m *Datatings) GetDesigner(permit string) (PermitModelPermitDesigner, error) {
	stmt := ""
	row := m.DB.QueryRow(stmt, permit)

	p := PermitModelPermitDesigner{}

	err := row.Scan(&p.Permit, &p.Designer, &p.DateStarted, &p.DateCompleted)

	if err != nil {
		return PermitModelPermitDesigner{}, err
	}
	return p, nil

}

//insert permit date received

func (m *Datatings) InsertDateReceived(p PermitModelPermitDateReceived) error {
	stmt := ""
	_, err := m.DB.Exec(stmt, p.Permit, p.DateReceived)
	if err != nil {
		return err
		log.Println(err)
	}

	return nil
}

/* TODO


insert design start
insert design finish


get permit date received
insert permit date submitted
get permit date submitted

*/
