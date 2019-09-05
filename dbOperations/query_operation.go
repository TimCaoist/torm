package torm

import (
	"torm/context"
	"torm/dbOperations/queryHandlers"
)

func Where(sql string, param interface{}, dbKey string) (interface{}, error) {
	c := &context.DBQueryContext{}
	c.Params = param
	c.QueryConfig = context.QueryConfig{}
	c.QueryConfig.DbKey = dbKey
	c.QueryConfig.Sql = sql
	queryHandler := queryHandlers.GetQueryHandler(c)
	return queryHandler.Query(c)
}

func SWhere(target interface{}, sql string, param interface{}, dbKey string) (interface{}, error) {
	c := &context.DBQueryContext{}
	c.Params = param
	c.QueryConfig = context.QueryConfig{}
	c.QueryConfig.DbKey = dbKey
	c.QueryConfig.Sql = sql
	queryHandler := queryHandlers.GetQueryHandler(c)
	return queryHandler.Query(c)
}
