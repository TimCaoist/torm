package valueSetter

import (
	"database/sql"
	"sync"
	"torm/context"
)

type MapValueSetter struct {
}

var mapValueSetterOnce sync.Once

var mapValueSetter *MapValueSetter

func GetMapValueSetterInstance() *MapValueSetter {
	mapValueSetterOnce.Do(func() {
		mapValueSetter = &MapValueSetter{}
	})

	return mapValueSetter
}

func (m MapValueSetter) Scan(config context.QueryConfig, contxt *context.DBQueryContext, rows *sql.Rows, cols []string) interface{} {
	lenCol := len(cols)
	values := make([][]byte, lenCol)
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, lenCol)
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	datas := make([]map[string]string, 0)
	for rows.Next() {
		_ = rows.Scan(scans...)
		data := make(map[string]string)
		for i := 0; i < lenCol; i++ {
			data[cols[i]] = string(values[i])
		}

		datas = append(datas, data)
	}

	return datas[:]
}
