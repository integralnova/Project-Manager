package models

import "fmt"

type Field struct {
	Name string
}

type Query struct {
	Table  string
	Fields map[string]Field
}

func (q Query) Select() string {
	//"`SELECT * FROM permitid ORDER BY id ASC`"
	fields := ""
	for _, i := range q.Fields {
		fields += q.Fields[i.Name].Name + ", "
		
	}
	stmt := fmt.Sprintf("SELECT %s FROM %s ", fields, q.Table)

	return stmt
}

func Insert() string {
	return ""
}

func Update() string {
	return ""
}

func Delete() string {
	return ""
}
