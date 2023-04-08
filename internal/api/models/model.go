package models

type Model struct {
	Table   Table
	Structs interface{}
}

type Table struct {
	Name    string
	Columns []string
}
