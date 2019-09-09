package valueSetter

import (
	"database/sql"
	"torm/context"
)

type IValueSetter interface {
	Scan(config *context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{}

	ScanOneRow(config *context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{}
}

func BuilderValueSetter(config *context.QueryConfig, c *context.DBQueryContext) IValueSetter {
	if config.Target == nil {
		return GetMapValueSetterInstance()
	}

	return GetStructValueSetterInstance()
}
