package queryHandlers

import (
	"torm/context"
)

type QueryHandler struct {
	Excuter IQueryHandler
}

type IQueryHandler interface {
	Query(context *context.DBQueryContext) (interface{}, error)
}

func (qh QueryHandler) Query(context *context.DBQueryContext) (interface{}, error) {
	return qh.Excuter.Query(context)
}

func GetQueryHandler(context *context.DBQueryContext) QueryHandler {
	q := QueryHandler{}
	q.Excuter = &SimpleQueryHandler{q}
	return q
}
