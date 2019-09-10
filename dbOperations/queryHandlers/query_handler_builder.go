package queryHandlers

import (
	"fmt"
	"torm/context"
)

const (
	Simple = iota
	Single_Row
)

var queryHandlers map[int]IQueryHandler

func Builder(config context.QueryConfig) (IQueryHandler, error) {
	handler, ok := queryHandlers[config.Type]
	if ok {
		return handler, nil
	}

	return nil, fmt.Errorf("Cann't found matching handler.")
}

func GetQueryHandler(target interface{}, sql string, param interface{}, dbKey string, queryType int) (*context.DBQueryContext, IQueryHandler, error) {
	c := &context.DBQueryContext{}
	c.Params = param
	c.QueryConfig = context.QueryConfig{}
	c.QueryConfig.DbKey = dbKey
	c.QueryConfig.Sql = sql
	c.QueryConfig.Type = queryType
	c.QueryConfig.Target = target
	queryHandler, err := Builder(c.QueryConfig)
	return c, queryHandler, err
}

func init() {
	queryHandlers = make(map[int]IQueryHandler)
	queryHandlers[Simple] = &SimpleQueryHandler{}
	queryHandlers[Single_Row] = &SingleRowQueryHandler{}
}
