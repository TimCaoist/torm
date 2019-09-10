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

type BatchDeleteHandler struct {
	UpdateHandler
}

func (qh BatchDeleteHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	sql, key, err := getDeleteInfo(qh.UpdateHandler, config)
	if err != nil {
		return err
	}

	if key == nil {
		sql = sql + config.UpdateModel.Filter
		config.Sql = sql
		return deleteDataByFilter(config, context)
	}

	return deleteDataByKey(sql, key, config, context)
}

func deleteDataByFilter(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	rVal := common.GetReflectIndirectValue(config.UpdateModel.Data)
	count := rVal.Len()
	if count == 0 {
		return fmt.Errorf("Datas is empty!")
	}

	for i := 0; i < count; i++ {
		context.Params = rVal.Index(i)
		err := sqlExcuter.UpdateOnTran(config, context)
		if err != nil {
			return err
		}
	}

	if config.IsOnTran == false {
		context.Commit()
	}

	return nil
}

func deleteDataByKey(prefiexSql string, key *dataMapping.MappingData, config *context.UpdateConfig, context *context.DBUpdateContext) error {
	rVal := common.GetReflectIndirectValue(config.UpdateModel.Data)
	count := rVal.Len()
	if count == 0 {
		return fmt.Errorf("Datas is empty!")
	}

	keyDatas := make([]string, count)
	for i := 0; i < count; i++ {
		rItem := rVal.Index(i)
		keyDatas[i] = fmt.Sprintf("'%v'", common.GetFieldValue(rItem, key.FieldName))
	}

	strBuffer := bytes.Buffer{}
	strBuffer.WriteString(prefiexSql)
	strBuffer.WriteString(key.DBName)
	strBuffer.WriteString(common.In)
	strBuffer.WriteString(common.Start)
	strBuffer.WriteString(strings.Join(keyDatas, common.Split))
	strBuffer.WriteString(common.End)

	context.Params = nil
	config.Sql = string(strBuffer.Bytes())
	return sqlExcuter.Update(config, context)
}
