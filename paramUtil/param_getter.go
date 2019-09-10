package paramUtil

import (
	"fmt"
	"reflect"
	"torm/common"
	"torm/context"
)

const (
	paramStr = "?"
)

type IParamGetter interface {
	Get(paramName string) interface{}
	GetArgs(matcheParams [][]int, sql string) (string, []interface{})
}

func GetParamGetter(context context.IDBContext) IParamGetter {
	var params = context.GetParams()
	if params == nil {
		return nil
	}

	switch params.(type) {
	case reflect.Value:
		return NewReflectParamGetter(params.(reflect.Value))
	case map[string]interface{}:
		v, _ := params.(map[string]interface{})
		return &MapParamGetter{Datas: v}
	case []interface{}:
		v, _ := params.([]interface{})
		return &DefaultParamGetter{Datas: v}
	}

	rValue := common.GetReflectIndirectValue(params)
	if rValue.Kind() == reflect.Struct {
		return NewReflectParamGetter(rValue)
	}

	return nil
}

func GetAllArgs(paramGetter IParamGetter, matchParams [][]int, sql string) (string, []interface{}) {
	lenResult := len(matchParams)
	values := make([]interface{}, lenResult)

	for i := lenResult - 1; i >= 0; i-- {
		v := matchParams[i]
		paramName := sql[v[0]+1 : v[1]-1]
		str1 := sql[0:v[0]]
		str2 := sql[v[1]-1 : len(sql)]
		sql = fmt.Sprintf("%s%s%s", str1, paramStr, str2)
		values[i] = paramGetter.Get(paramName)
	}

	return sql, values[:]
}
