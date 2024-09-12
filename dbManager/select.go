package dbmanager

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

//TODO:
