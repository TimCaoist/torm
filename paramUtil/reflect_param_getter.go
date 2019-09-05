package paramUtil

import (
	"reflect"
	"strings"
	"torm/common"
)

type ReflectParamGetter struct {
	ReflectValue reflect.Value
	Values       map[string]interface{}
}

const (
	paramSplit = "."
)

func (reflectParamGetter ReflectParamGetter) Get(paramName string) interface{} {
	value, ok1 := reflectParamGetter.Values[paramName]
	if ok1 {
		return value
	}

	reflectParamGetter.Values[paramName] = GetReflectValue(paramName, reflectParamGetter.ReflectValue)
	return reflectParamGetter.Values[paramName]
}

func (reflectParamGetter ReflectParamGetter) GetArgs(matcheParams [][]int, sql string) (string, []interface{}) {
	return GetAllArgs(reflectParamGetter, matcheParams, sql)
}

func GetReflectValue(paramName string, rValue reflect.Value) interface{} {
	rValue = common.Indirect(rValue)
	fileds := strings.Split(paramName, paramSplit)
	field := rValue.FieldByName(fileds[0])
	if field.IsValid() == false {
		return nil
	}

	if len(fileds) == 1 {
		return field.Interface()
	}

	return GetReflectValue(strings.Join(fileds[1:], ""), reflect.ValueOf(field.Interface()))
}
