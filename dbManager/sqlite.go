package dbmanager

import (
	"database/sql"

	models "github.com/integralnova/Project-Manager/internal"
)

type PermitModel struct {
	DB *sql.DB
}

func (m *PermitModel) Insert(title, content string) error {
	stmt := `INSERT INTO posts (title, content, createdAt)
	VALUES (?, ?, datetime('now'))`
	_, err := m.DB.Exec(stmt, title, content)
	return err
}

func (m *PermitModel) All() ([]models.Permits, error) {
	stmt := `SELECT id, title, content, createdAt FROM posts ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	permits := []models.Permits{}
	for rows.Next() {
		p := models.Permits{}
		err := rows.Scan(&p.ID, &p.PermitID, &p.DateReceived, &p.Designer)
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
