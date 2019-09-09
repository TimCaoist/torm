package updateHandlers

import (
	"bytes"
	"fmt"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
	"torm/dbOperations/queryHandlers"
	"torm/sqlExcuter"
)

type BatchInsertHandler struct {
	UpdateHandler
}

func QueryMaxId(key *dataMapping.MappingData, tableName string, dbKey string) int64 {
	sql := fmt.Sprintf("select %v from %v order by Id desc limit 0, 1 ", key.DBName, tableName)
	var id int64
	c, handler, err := queryHandlers.GetQueryHandler(&id, sql, nil, dbKey, 2)
	if err != nil {
		return id
	}

	handler.Query(&c.QueryConfig, c)
	return id
}

func (qh BatchInsertHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	updateMappings, tableName, key, err := GetBacthUpdateInfo(qh.UpdateHandler, config, context, false)
	if err != nil {
		return err
	}

	updateModel := config.UpdateModel
	rVal := common.GetReflectIndirectValue(updateModel.Data)
	count := rVal.Len()
	if count == 0 {
		return fmt.Errorf("Cann't insert empty data.")
	}

	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.InsertInto)
	strBuffer.WriteString(tableName)
	strBuffer.WriteString(common.Start)

	hasKey := false
	len := len(*updateMappings)
	for i, col := range *updateMappings {
		strBuffer.WriteString(col.DBName)

		if col.DBName == key.DBName {
			hasKey = true
		}

		if i != len-1 {
			strBuffer.WriteString(common.Split)
		}
	}

	strBuffer.WriteString(common.End)
	strBuffer.WriteString(common.Values)

	if hasKey {
		firstVal := rVal.Index(0)
		JudgeRequireReturnId(config, context, key, firstVal)
	}

	var maxId int64
	if config.RequireId {
		maxId = QueryMaxId(key, tableName, config.DbKey)
	}

	lenData := rVal.Len()
	for i := 0; i < lenData; i++ {
		itemVal := rVal.Index(i)
		strBuffer.WriteString(common.Start)
		for ci, col := range *updateMappings {
			fieldValue := itemVal.FieldByName(col.FieldName)
			if col.DBName == key.DBName && config.RequireId {
				fieldValue.SetInt(maxId + int64(i) + 1)
			}

			interfaceVal := fieldValue.Interface()
			strBuffer.WriteString(common.Split1)
			strBuffer.WriteString(fmt.Sprintf("%v", interfaceVal))
			strBuffer.WriteString(common.Split1)
			if ci != len-1 {
				strBuffer.WriteString(common.Split)
			}
		}

		strBuffer.WriteString(common.End)
		if i != lenData-1 {
			strBuffer.WriteString(common.Split)
		}
	}

	config.Sql = string(strBuffer.Bytes())
	return sqlExcuter.Update(config, context)
}
