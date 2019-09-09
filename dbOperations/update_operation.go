package torm

import (
	"torm/common"
	"torm/context"
	"torm/dbOperations/updateHandlers"
)

func Insert(data interface{}, dbKey string, fields []string) error {
	isSlice := common.IsSlice(data)
	if isSlice {
		return BatchInsert(data, dbKey, fields)
	}

	updateModel := context.UpdateModel{}
	updateModel.Data = data
	return UpdateByModel(updateModel, dbKey, updateHandlers.Single_Insert)
}

func UpdateByModel(model context.UpdateModel, dbKey string, excuteType int) error {
	c := context.NewDBUpdateContext()
	c.Params = model.Data
	c.UpdateConfig = context.UpdateConfig{}
	c.UpdateConfig.UpdateModel = model
	c.UpdateConfig.DbKey = dbKey
	c.UpdateConfig.Type = excuteType
	queryHandler := updateHandlers.GetUpdateHandler(c)
	return queryHandler.Update(c)
}

func Update(data interface{}, dbKey string, fields []string, filter string) error {
	isSlice := common.IsSlice(data)
	if isSlice {
		return BatchUpdate(data, dbKey, fields, filter)
	}

	updateModel := context.UpdateModel{}
	updateModel.Data = data
	updateModel.Fields = fields
	updateModel.Filter = filter
	return UpdateByModel(updateModel, dbKey, updateHandlers.Single_Update)
}

func BatchUpdate(datas interface{}, dbKey string, fields []string, filter string) error {
	updateModel := context.UpdateModel{}
	updateModel.Data = datas
	updateModel.Fields = fields
	updateModel.Filter = filter
	updateType := updateHandlers.Batch_Update
	if filter != common.Empty {
		updateType = updateHandlers.Batch_Update_Filter
	}

	return UpdateByModel(updateModel, dbKey, updateType)
}

func BatchInsert(datas interface{}, dbKey string, fields []string) error {
	updateModel := context.UpdateModel{}
	updateModel.Data = datas
	updateModel.Fields = fields
	updateType := updateHandlers.Batch_Inert
	return UpdateByModel(updateModel, dbKey, updateType)
}

func UpdateOnTranByModel(model context.UpdateModel, dbKey string, excuteType int, c *context.DBUpdateContext) (*context.DBUpdateContext, error) {
	if c == nil {
		c = context.NewDBUpdateContext()
	}

	c.Params = model.Data
	c.UpdateConfig = context.UpdateConfig{}
	c.UpdateConfig.UpdateModel = model
	c.UpdateConfig.DbKey = dbKey
	c.UpdateConfig.Type = excuteType
	c.UpdateConfig.IsOnTran = true
	queryHandler := updateHandlers.GetUpdateHandler(c)
	return c, queryHandler.Update(c)
}

func UpdateOnTran(data interface{}, dbKey string, excuteType int, c *context.DBUpdateContext) (*context.DBUpdateContext, error) {
	updateModel := context.UpdateModel{}
	updateModel.Data = data
	return UpdateOnTranByModel(updateModel, dbKey, excuteType, c)
}
