package torm

import (
	"torm/context"
	"torm/dbOperations/queryHandlers"
)

func Where(sql string, param interface{}, dbKey string) (interface{}, error) {
	c := &context.DBQueryContext{}
	c.Params = param
	c.QueryConfig = context.QueryConfig{DbKey: dbKey, Sql: sql}
	queryHandler := queryHandlers.GetQueryHandler(c)
	return queryHandler.Query(c)
}

func SWhere(target interface{}, sql string, param interface{}, dbKey string) (interface{}, error) {
	c := &context.DBQueryContext{}
	c.Params = param
	c.QueryConfig = context.QueryConfig{DbKey: dbKey, Sql: sql, Target: target}
	queryHandler := queryHandlers.GetQueryHandler(c)
	return queryHandler.Query(c)
}
