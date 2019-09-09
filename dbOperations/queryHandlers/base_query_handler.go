package queryHandlers

import (
	"torm/context"
)

type QueryHandler struct {
	Excuter IQueryHandler
}

type IQueryHandler interface {
	Query(queryConfig *context.QueryConfig, context *context.DBQueryContext) (interface{}, error)
}

func (qh QueryHandler) Query(context *context.DBQueryContext) (interface{}, error) {
	return qh.Excuter.Query(&context.QueryConfig, context)
}
