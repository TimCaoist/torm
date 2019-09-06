package updateHandlers

import (
	"bytes"
	"strings"
	"torm/common"
	"torm/context"
	"torm/sqlExcuter"
)

type SingleInsertHandler struct {
	UpdateHandler
}

func (qh SingleInsertHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	if config.Sql != common.Empty {
		return sqlExcuter.Update(context.UpdateConfig, context)
	}

	tableName, mappingDatas := qh.GetStructInfo(config)

	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.InsertInto)
	strBuffer.WriteString(tableName)
	cols := []string{}
	fields := []string{}
	for _, v := range mappingDatas {
		if v.Ingore == true {
			continue
		}

		cols = append(cols, v.DBName)
		fields = append(fields, common.ParamStart+v.FieldName)
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

	context.UpdateConfig.Sql = string(strBuffer.Bytes())
	return sqlExcuter.Update(context.UpdateConfig, context)
}
