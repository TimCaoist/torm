package updateHandlers

import (
	"fmt"
	"torm/common"
	"torm/context"
	"torm/sqlExcuter"
)

type BatchUpdateFilterHandler struct {
	UpdateHandler
}

func (qh BatchUpdateFilterHandler) Update(config *context.UpdateConfig, context *context.DBUpdateContext) error {
	updateModel := config.UpdateModel
	rVal := common.GetReflectIndirectValue(updateModel.Data)
	count := rVal.Len()

	if count == 0 {
		return fmt.Errorf("Datas is empty!")
	}

	updateMappings, tableName, key, err := GetBacthUpdateInfo(qh.UpdateHandler, config, context, true)
	if err != nil {
		return err
	}

	sql := config.Sql
	if sql == common.Empty {
		sql = BuilderUpdateSql(tableName, *key, config.UpdateModel, updateMappings)
	}

	config.Sql = sql
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
