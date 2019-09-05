package valueSetter

import (
	"database/sql"
	"reflect"
	"sync"
	"torm/common"
	"torm/context"
	"torm/dataMapping"
)

type StructValueSetter struct {
}

var structValueSetterOnce sync.Once

var structValueSetter *StructValueSetter

func GetStructValueSetterInstance() *StructValueSetter {
	structValueSetterOnce.Do(func() {
		structValueSetter = &StructValueSetter{}
	})

	return structValueSetter
}

func (s StructValueSetter) Scan(config context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{} {
	rVal := common.Indirect(reflect.ValueOf(config.Target))
	isSlice := rVal.Kind() == reflect.Slice
	if isSlice {
		return s.ScanArray(rVal, config, contxt, rows, cols)
	}

	return s.ScanSingle(rVal, config, contxt, rows, cols)
}

func (s StructValueSetter) ScanArray(rVal reflect.Value, config context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{} {
	elementType := rVal.Type().Elem()
	mappingData := dataMapping.GetTypeMapping(elementType)
	datas := make([]reflect.Value, 0)
	for rows.Next() {
		data := reflect.New(elementType).Elem()
		s.ScanRow(rows, cols, data, mappingData)
		datas = append(datas, data)
	}

	v := reflect.Append(rVal, datas...)
	rVal.Set(v)
	return nil
}

func (s StructValueSetter) ScanSingle(rVal reflect.Value, config context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{} {
	mappingData := dataMapping.GetTypeMapping(rVal.Type())
	for rows.Next() {
		s.ScanRow(rows, cols, rVal, mappingData)
		break
	}

	return nil
}

func (s StructValueSetter) ScanRow(rows *sql.Rows, columns []string, data reflect.Value, mappingDatas []dataMapping.MappingData) {
	lenCol := len(columns)
	values := make([]interface{}, lenCol)
	var ingore interface{}
	for a := 0; a < lenCol; a++ {
		mappingData, ok := dataMapping.GetMatchMapingData(columns[a], mappingDatas)
		if ok == false || mappingData.Ingore {
			values[a] = &ingore
			continue
		}

		field := data.FieldByName(mappingData.FieldName)
		if field.CanAddr() == false {
			values[a] = &ingore
			continue
		}

		values[a] = field.Addr().Interface()
	}

	rows.Scan(values...)
}
