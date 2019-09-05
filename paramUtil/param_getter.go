package paramUtil

import (
	"torm/context"
)

type IParamGetter interface {
	Get(paramName string) interface{}
}

type MapParamGetter struct {
	Datas   map[string]interface{}
	Context context.DBContext
}

func GetParamGetter(context context.IDBContext) IParamGetter {
	var params = context.GetParams()
	switch params.(type) {
	case map[string]interface{}:
		v, _ := params.(map[string]interface{})
		return &MapParamGetter{Datas: v}
	}

	return nil
}

func (mapParamGetter MapParamGetter) Get(paramName string) interface{} {
	return mapParamGetter.Datas[paramName]
}
