package updateHandlers

import (
	"bytes"
	"fmt"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
	"torm/sqlExcuter"
)

type SingleDeleteHandler struct {
	UpdateHandler
}

func (qh SingleDeleteHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	sql, key, err := getDeleteInfo(qh.UpdateHandler, config)
	if err != nil {
		return err
	}

	if key == nil {
		sql = sql + config.UpdateModel.Filter
		config.Sql = sql
		return sqlExcuter.Update(config, context)
	}

	sql = fmt.Sprintf("%v%v = @%v ", sql, key.DBName, key.FieldName)
	config.Sql = sql
	return sqlExcuter.Update(config, context)
}

func buildPrefixDeleteSql(tableName string) string {
	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(common.Delete)
	strBuffer.WriteString(tableName)
	strBuffer.WriteString(common.Where)

	prefixSql := string(strBuffer.Bytes())
	return prefixSql
}

func getDeleteInfo(qh UpdateHandler, config *context.UpdateConfig) (string, *dataMapping.MappingData, error) {
	tableName, mappingDatas := qh.GetStructInfo(config)
	key, isFound := qh.GetKey(*mappingDatas)
	updateModel := config.UpdateModel
	if updateModel.Filter != common.Empty {
		return buildPrefixDeleteSql(tableName), nil, nil
	}

	if isFound == false {
		return common.Empty, nil, fmt.Errorf("Please setting the filter value when without key field")
	}

	prefixSql := buildPrefixDeleteSql(tableName)
	return prefixSql, key, nil
}
