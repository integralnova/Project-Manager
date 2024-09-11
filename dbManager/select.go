package dbmanager

import (
	"database/sql"
	"fmt"
)

func (m *PermitModel) CreatePermit(title string, content string) error {
	stmt := `INSERT INTO permits (title, content, createdAt)
	VALUES (?, ?, datetime('now'))`
	_, err := m.DB.Exec(stmt, title, content)
	return err
}

func (m *PermitModel) GetAllPermits() ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitByID(id int) (*Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)

	var permit Permit
	err := row.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no permit found with id %d", id)
		}
		return nil, err
	}
	return &permit, nil
}

func (m *PermitModel) UpdatePermit(id int, title string, content string) error {
	stmt := `UPDATE permits SET title = ?, content = ? WHERE id = ?`
	_, err := m.DB.Exec(stmt, title, content, id)
	return err
}

func (m *PermitModel) DeletePermit(id int) error {
	stmt := `DELETE FROM permits WHERE id = ?`
	_, err := m.DB.Exec(stmt, id)
	return err
}

func (m *PermitModel) GetPermitsByStatus(status string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE status = ?`
	rows, err := m.DB.Query(stmt, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByDesigner(designer string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE designer = ?`
	rows, err := m.DB.Query(stmt, designer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByOwner(owner string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE owner = ?`
	rows, err := m.DB.Query(stmt, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByDateReceived(dateReceived string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE dateReceived = ?`
	rows, err := m.DB.Query(stmt, dateReceived)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByDateApproved(dateApproved string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE dateApproved = ?`
	rows, err := m.DB.Query(stmt, dateApproved)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByDesignStartDate(designStartDate string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE designStartDate = ?`
	rows, err := m.DB.Query(stmt, designStartDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByDesignFinishDate(designFinishDate string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE designFinishDate = ?`
	rows, err := m.DB.Query(stmt, designFinishDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

func (m *PermitModel) GetPermitsByPermitType(permitType string) ([]Permit, error) {
	stmt := `SELECT id, title, content, createdAt FROM permits WHERE permitType = ?`
	rows, err := m.DB.Query(stmt, permitType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var permit Permit
		err := rows.Scan(&permit.ID, &permit.Title, &permit.Content, &permit.CreatedAt)
		if err != nil {
			return nil, err
		}
		permits = append(permits, permit)
	}
	return permits, nil
}

type Permit struct {
	ID               int
	Title            string
	Content          string
	CreatedAt        string
	Status           string
	Designer         string
	Owner            string
	DateReceived     string
	DateApproved     string
	DesignStartDate  string
	DesignFinishDate string
	PermitType       string
}
