package context

type UpdateModel struct {
	Data      interface{}
	TableName string
	Fields    []string
	Filter    string
	Sql       string
}
