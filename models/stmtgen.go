package models

import "fmt"

type Field struct {
	Name string
}

type Query struct {
	Table  string
	Fields map[string]Field
}

//TODO Actually learn sql

func (q Query) Select() string {
	//"`SELECT * FROM permitid ORDER BY id ASC`"
	fields := ""
	for _, i := range q.Fields {
		fields += q.Fields[i.Name].Name + ", "

	}
	stmt := fmt.Sprintf("SELECT %s FROM %s ", fields, q.Table)

	return stmt
}
func (m *Datatings) GenericSelect(query string, dest *[]PermitsModel) error {
	rows, err := m.DB.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result PermitsModel
	for rows.Next() {
		// Scan the row data into result
		if err := rows.Scan(&result.ID, &result.PermitID, &result.CompanyName, &result.Designer, &result.DateReceived, &result.DateDue, &result.PermitStatus, &result.Designer); err != nil {
			return err
		}
		// Append the result to the slice
		*dest = append(*dest, result)
	}

	return rows.Err() // Check for any errors encountered during iteration
}