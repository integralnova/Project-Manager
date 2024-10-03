package models

import "fmt"

// How to deal with duplicates.
// Find at creation. use UPSERT
// Might have duplicate permit names
// Insert permit
func (m *Datatings) InsertPermit(permit PermitsModel) error {
	fmt.Println("Inserting permits")
	stmt := `INSERT INTO permits (permitID, companyName, reference, dateReceived, dateDue, permitStatus, designer)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, permit.PermitID, permit.CompanyName, permit.Reference, permit.DateReceived, permit.DateDue, permit.PermitStatus, permit.Designer)
	if err != nil {
		return err
	}

	return nil

}


// All permits
// NEEDS rewrite
func (m *Datatings) Getpermits() ([]PermitsViewModel, error) {
	stmt := `SELECT * FROM permits ORDER BY id ASC`
	var permit []PermitsModel
	err := m.GenericSelect(stmt, &permit)
	if err != nil {
		return nil, err
	}

	p := make([]PermitsViewModel, len(permit))
	for i := range permit {
		p[i] = TranslatePermit(permit[i])
	}
	return p, nil
}


