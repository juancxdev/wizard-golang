package main

// Model estructura que se generará
type Model struct {
	ID            string
	Name          string
	Table         string
	Fields        []Field
	PackageRoutes map[string]string
	LogicDelete   string
}

// Field structura de un tipo de campo del modelo
type Field struct {
	Name    string
	Type    string
	NotNull string
	Len     int
}
