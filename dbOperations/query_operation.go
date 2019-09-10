package torm

import (
	"torm/dbOperations/queryHandlers"
)

func Raw(sql string, param interface{}, dbKey string) (interface{}, error) {
	c, handler, err := queryHandlers.GetQueryHandler(nil, sql, param, dbKey, queryHandlers.Simple)
	if err != nil {
		return nil, err
	}

	return handler.Query(&c.QueryConfig, c)
}

func Single(target interface{}, sql string, param interface{}, dbKey string) (interface{}, error) {
	c, handler, err := queryHandlers.GetQueryHandler(target, sql, param, dbKey, queryHandlers.Single_Row)
	if err != nil {
		return nil, err
	}

	return handler.Query(&c.QueryConfig, c)
}

func SRaw(target interface{}, sql string, param interface{}, dbKey string) (interface{}, error) {
	c, handler, err := queryHandlers.GetQueryHandler(target, sql, param, dbKey, queryHandlers.Simple)
	if err != nil {
		return nil, err
	}

	return handler.Query(&c.QueryConfig, c)
}
