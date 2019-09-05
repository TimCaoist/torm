package queryHandlers

import (
	"torm/context"
	"torm/sqlExcuter"
)

type SimpleQueryHandler struct {
	QueryHandler
}

func (qh SimpleQueryHandler) Query(context *context.DBQueryContext) (interface{}, error) {
	return sqlExcuter.Query(context.QueryConfig, context)
}
