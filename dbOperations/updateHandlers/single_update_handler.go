package updateHandlers

import (
	"bytes"
	"fmt"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
	"torm/sqlExcuter"
)

type SingleUpdateHandler struct {
	UpdateHandler
}

func BuilderUpdateSql(tableName string, key dataMapping.MappingData, updateModel context.UpdateModel, mappingDatas *[]dataMapping.MappingData) string {
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.Update)
	strBuffer.WriteString(tableName)
	strBuffer.WriteString(common.Set)
	updateFields := updateModel.Fields

	if len(updateFields) == 0 {
		lenM := len(*mappingDatas) - 1
		for i, v := range *mappingDatas {
			if v.FieldName == key.FieldName {
				continue
			}

			BuildUpdateCol(&strBuffer, &v, i == lenM)
		}
	} else {
		lenM := len(updateFields) - 1
		for i, v := range updateFields {
			m, ok := dataMapping.GetMatchMapingData(v, *mappingDatas)
			if !ok {
				continue
			}

			BuildUpdateCol(&strBuffer, &m, i == lenM)
		}
	}

	strBuffer.WriteString(common.Where)
	if updateModel.Filter != common.Empty {
		strBuffer.WriteString(updateModel.Filter)
	} else {
		BuildUpdateCol(&strBuffer, &key, true)
	}

	strBuffer.WriteString(common.WhiteSpace)
	return string(strBuffer.Bytes())
}

func (qh SingleUpdateHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	if config.Sql != common.Empty {
		return sqlExcuter.Update(config, context)
	}

	updateModel := config.UpdateModel
	tableName, mappingDatas := qh.GetStructInfo(config)
	key, isFound := qh.GetKey(*mappingDatas)
	if isFound == false && updateModel.Filter == common.Empty {
		return fmt.Errorf("Please setting the filter value when without key field")
	}

	config.Sql = BuilderUpdateSql(tableName, *key, config.UpdateModel, mappingDatas)
	return sqlExcuter.Update(config, context)
}

func BuildUpdateCol(strBuffer *bytes.Buffer, mappingData *dataMapping.MappingData, isLast bool) {
	strBuffer.WriteString(mappingData.DBName)
	strBuffer.WriteString(common.Equlas)
	strBuffer.WriteString(common.ParamStart)
	strBuffer.WriteString(mappingData.FieldName)

	if isLast == false {
		strBuffer.WriteString(common.Split)
	}
}
