package paramUtil

type DefaultParamGetter struct {
	Datas []interface{}
}

func (paramGetter DefaultParamGetter) Get(paramName string) interface{} {
	return nil
}

func (paramGetter DefaultParamGetter) GetArgs(matcheParams [][]int, sql string) (string, []interface{}) {
	return sql, paramGetter.Datas
}
