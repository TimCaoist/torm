package updateHandlers

import (
	"reflect"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
)

type UpdateHandler struct {
	Excuter *IUpdateHandler
}

var emptyMappingData dataMapping.MappingData

type IUpdateHandler interface {
	Update(config *context.UpdateConfig, context *context.DBUpdateContext) error
}

func (qh *UpdateHandler) Update(context *context.DBUpdateContext) error {
	excuter := *qh.Excuter
	return excuter.Update(&context.UpdateConfig, context)
}

func (u UpdateHandler) GetStructInfo(config *context.UpdateConfig) (string, []dataMapping.MappingData) {
	updateModel := config.UpdateModel
	tableName := updateModel.TableName
	rType := common.GetReflectIndirectType(updateModel.Data)
	if rType.Kind() == reflect.Slice {
		rValue := common.GetReflectIndirectValue(updateModel.Data)
		rType = rValue.Type().Elem()
	}

	if tableName == common.Empty {
		tableName = rType.Name()
	}

	return tableName, dataMapping.GetTypeMapping(rType)[:]
}

func (u UpdateHandler) GetKey(mappingDatas []dataMapping.MappingData) (*dataMapping.MappingData, bool) {
	for i, v := range mappingDatas {
		if v.IsKey {
			return &mappingDatas[i], true
		}
	}

	return &emptyMappingData, false
}

func GetUpdateHandler(context *context.DBUpdateContext) *UpdateHandler {
	q := &UpdateHandler{}
	handler, err := Builder(context.UpdateConfig)
	if err != nil {
		panic(err)
	}

	q.Excuter = handler
	return q
}
