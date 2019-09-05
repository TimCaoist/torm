package torm

import (
	"torm/context"
	"torm/dbOperations/updateHandlers"
)

type UpdateModel struct {
	Data      interface{}
	TableName string
}

func Insert(model UpdateModel, dbKey string) error {
	c := &context.DBUpdateContext{}
	c.Params = model
	c.UpdateConfig = context.UpdateConfig{}
	c.UpdateConfig.DbKey = dbKey
	queryHandler := updateHandlers.GetUpdateHandler(c)
	return queryHandler.Update(c)
}
