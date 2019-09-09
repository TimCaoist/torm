package queryHandlers

import (
	"torm/context"
	"torm/sqlExcuter"
)

type SimpleQueryHandler struct {
	QueryHandler
}

func (qh SimpleQueryHandler) Query(queryConfig *context.QueryConfig, context *context.DBQueryContext) (interface{}, error) {
	return sqlExcuter.Query(queryConfig, context)
}
