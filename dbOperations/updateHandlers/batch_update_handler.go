package updateHandlers

import (
	"bytes"
	"fmt"
	"strings"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
	"torm/sqlExcuter"
)

type BatchUpdateHandler struct {
	UpdateHandler
}

const (
	when_then_formatter = " WHEN '%v' THEN '%v' "
	case_formatter      = " `%v` = CASE `%v` "
	val_formatter       = "%v"
)

func GetBacthUpdateInfo(qh UpdateHandler, config *context.UpdateConfig, context *context.DBUpdateContext, ingoreKey bool) (*[]dataMapping.MappingData, string, *dataMapping.MappingData, error) {
	tableName, mappingDatas := qh.GetStructInfo(config)
	key, isFound := qh.GetKey(*mappingDatas)
	if isFound == false {
		return nil, common.Empty, key, fmt.Errorf("Current handler only support Key filter.")
	}

	updateModel := config.UpdateModel
	updateMappings := []dataMapping.MappingData{}
	if len(updateModel.Fields) == 0 {
		for _, v := range *mappingDatas {
			if ingoreKey && v.FieldName == key.FieldName {
				continue
			}

			updateMappings = append(updateMappings, v)
		}
	} else {
		for _, v := range updateModel.Fields {
			m, ok := dataMapping.GetMatchMapingData(v, *mappingDatas)
			if !ok {
				continue
			}

			updateMappings = append(updateMappings, m)
		}
	}

	return &updateMappings, tableName, key, nil
}

func (qh BatchUpdateHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	updateMappings, tableName, key, err := GetBacthUpdateInfo(qh.UpdateHandler, config, context, true)
	if err != nil {
		return err
	}

	updateModel := config.UpdateModel
	rVal := common.GetReflectIndirectValue(updateModel.Data)
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.Update)
	strBuffer.WriteString(tableName)
	strBuffer.WriteString(common.Set)

	count := rVal.Len()
	colCount := len(*updateMappings)
	colMappings := make(map[string]*bytes.Buffer, colCount)

	ids := make([]string, count)
	for idx := 0; idx < count; idx++ {
		dVal := common.Indirect(rVal.Index(idx))
		keyField := dVal.FieldByName(key.FieldName)
		keyValue := fmt.Sprintf(val_formatter, keyField.Interface())
		ids[idx] = keyValue

		for _, m := range *updateMappings {
			colBuffer, ok := colMappings[m.DBName]
			if ok == false {
				colBuffer = &bytes.Buffer{}
				colMappings[m.DBName] = colBuffer
			}

			field := dVal.FieldByName(m.FieldName)
			colBuffer.WriteString(fmt.Sprintf(when_then_formatter, keyValue, field.Interface()))
		}
	}

	colIndex := 1
	for k, v := range colMappings {
		strBuffer.WriteString(fmt.Sprintf(case_formatter, k, key.DBName))
		strBuffer.WriteString(string(v.Bytes()))
		strBuffer.WriteString(common.CaseEnd)
		if colIndex < colCount {
			strBuffer.WriteString(common.Split)
		}

		colIndex++
	}

	strBuffer.WriteString(common.Where)
	strBuffer.WriteString(key.DBName)
	strBuffer.WriteString(common.In)
	strBuffer.WriteString(common.Start)
	strBuffer.WriteString(strings.Join(ids, common.Split))
	strBuffer.WriteString(common.End)

	config.Sql = string(strBuffer.Bytes())
	return sqlExcuter.Update(config, context)
}
