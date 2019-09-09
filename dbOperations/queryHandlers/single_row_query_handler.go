package queryHandlers

import (
	"torm/context"
	"torm/sqlExcuter"
)

type SingleRowQueryHandler struct {
	QueryHandler
}

func (qh SingleRowQueryHandler) Query(queryConfig *context.QueryConfig, context *context.DBQueryContext) (interface{}, error) {
	queryConfig.OnlyOneRow = true
	return sqlExcuter.Query(queryConfig, context)
}
