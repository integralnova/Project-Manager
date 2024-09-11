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

// this could ugly. Need to figure out if i can some how break all permits in to pages
// Look into limit offset
func (m *PermitModel) Getpermits() ([]models.PermitsModel, error) {
	stmt := `SELECT id, title, content, createdAt FROM posts ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	permits := []models.PermitsModel{}

	for rows.Next() {
		p := models.PermitsModel{}
		err := rows.Scan(&p.ID, &p.PermitID, &p.DateReceived, &p.Owner)
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
