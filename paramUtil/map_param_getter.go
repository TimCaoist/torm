package paramUtil

type MapParamGetter struct {
	Datas map[string]interface{}
}

func (mapParamGetter MapParamGetter) Get(paramName string) interface{} {
	return mapParamGetter.Datas[paramName]
}

func (mapParamGetter MapParamGetter) GetArgs(matcheParams [][]int, sql string) (string, []interface{}) {
	return GetAllArgs(mapParamGetter, matcheParams, sql)
}
