package torm

import (
	"torm/context"
	"torm/dbOperations/updateHandlers"
)

func Insert(model context.UpdateModel, dbKey string) error {
	c := &context.DBUpdateContext{}
	c.UpdateModel = model
	c.Params = model.Data
	c.UpdateConfig = context.UpdateConfig{}
	c.UpdateConfig.UpdateModel = model
	c.UpdateConfig.DbKey = dbKey
	c.UpdateConfig.Type = 1
	queryHandler := updateHandlers.GetUpdateHandler(c)
	return queryHandler.Update(c)
}
