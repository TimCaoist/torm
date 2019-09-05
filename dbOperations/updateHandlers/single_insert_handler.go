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

const (
	split      = ", "
	paramStart = "@"
	insertInto = "INSERT INTO "
	start      = "("
	end        = ")"
	values     = " VALUES "
	empty      = ""
)

type SingleInsertHandler struct {
	UpdateHandler
}

func (qh SingleInsertHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	if config.Sql != empty {
		return sqlExcuter.Update(context.UpdateConfig, context)
	}

	updateModel := config.UpdateModel
	tableName := updateModel.TableName
	rType := common.IndirectType(reflect.TypeOf(updateModel.Data))
	if tableName == empty {
		tableName = rType.Name()
	}

	mappingDatas := dataMapping.GetTypeMapping(rType)
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(insertInto)
	strBuffer.WriteString(tableName)
	cols := []string{}
	fields := []string{}
	for _, v := range mappingDatas {
		if v.Ingore == true {
			continue
		}

		cols = append(cols, v.DBName)
		fields = append(fields, paramStart+v.FieldName)
	}

	//Create Columns
	strBuffer.WriteString(start)
	strBuffer.WriteString(strings.Join(cols, split))
	strBuffer.WriteString(end)
	strBuffer.WriteString(values)

	//Create Fields
	strBuffer.WriteString(start)
	strBuffer.WriteString(strings.Join(fields, split))
	strBuffer.WriteString(end)

	context.UpdateConfig.Sql = string(strBuffer.Bytes())
	return sqlExcuter.Update(context.UpdateConfig, context)
}
