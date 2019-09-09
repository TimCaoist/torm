package updateHandlers

import (
	"bytes"
	"reflect"
	"strings"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
	"torm/sqlExcuter"
)

type SingleInsertHandler struct {
	UpdateHandler
}

const (
	emptyId = 0
)

func JudgeRequireReturnId(config *context.UpdateConfig, context *context.DBUpdateContext, key *dataMapping.MappingData, rVal reflect.Value) {
	field := common.Indirect(rVal.FieldByName(key.FieldName))
	fieldValue := field.Interface()
	switch field.Kind() {
	case reflect.Int64:
		config.RequireId = fieldValue.(int64) == emptyId
	case reflect.Int32:
		config.RequireId = fieldValue.(int32) == emptyId
	case reflect.Int:
		config.RequireId = fieldValue.(int) == emptyId
	}

	config.KeyField = field
}

func BuilderInsertSql(tableName string, config *context.UpdateConfig, mappingDatas *[]dataMapping.MappingData) string {
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.InsertInto)
	strBuffer.WriteString(tableName)
	cols := []string{}
	fields := []string{}
	insertFileds := config.UpdateModel.Fields

	if len(insertFileds) == 0 {
		for _, v := range *mappingDatas {
			if v.Ingore == true {
				continue
			}

			cols = append(cols, v.DBName)
			fields = append(fields, common.ParamStart+v.FieldName)
		}
	} else {
		for _, v := range insertFileds {
			m, ok := dataMapping.GetMatchMapingData(v, *mappingDatas)
			if !ok {
				continue
			}

			cols = append(cols, m.DBName)
			fields = append(fields, common.ParamStart+m.FieldName)
		}
	}

	//Create Columns
	strBuffer.WriteString(common.Start)
	strBuffer.WriteString(strings.Join(cols, common.Split))
	strBuffer.WriteString(common.End)
	strBuffer.WriteString(common.Values)

	//Create Fields
	strBuffer.WriteString(common.Start)
	strBuffer.WriteString(strings.Join(fields, common.Split))
	strBuffer.WriteString(common.End)
	return string(strBuffer.Bytes())
}

func (qh SingleInsertHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	tableName, mappingDatas := qh.GetStructInfo(config)
	key, isFound := qh.GetKey(*mappingDatas)
	if isFound {
		JudgeRequireReturnId(config, context, key, common.GetReflectIndirectValue(config.UpdateModel.Data))
	}

	if config.Sql != common.Empty {
		return qh.DoUpdate(config, context)
	}

	context.UpdateConfig.Sql = BuilderInsertSql(tableName, config, mappingDatas)
	return qh.DoUpdate(config, context)
}

func (qh SingleInsertHandler) DoUpdate(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	err := sqlExcuter.Update(config, context)
	if err != nil || config.RequireId == false {
		return err
	}

	field := config.KeyField
	field.SetInt(config.Id)
	return nil
}
