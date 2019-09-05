package parser

import (
	"errors"
	"fmt"
	"regexp"
	"torm/context"
	"torm/paramUtil"
)

const (
	paramReg  = "@.*?[, ]"
	paramStr  = "?"
	whieSpace = " "
)

func Parser(sql string, context context.IDBContext) (string, []interface{}, error) {
	reg := regexp.MustCompile(paramReg)
	//返回匹配到的结果
	result := reg.FindAllStringIndex(sql, -1)
	if len(result) == 0 {
		return sql, nil, nil
	}

	paramGetter := paramUtil.GetParamGetter(context)
	if paramGetter == nil {
		err := errors.New("Not Support")
		return sql, nil, err
	}

	lenResult := len(result)
	values := make([]interface{}, lenResult)

	for i := lenResult - 1; i >= 0; i-- {
		v := result[i]
		paramName := sql[v[0]+1 : v[1]-1]
		str1 := sql[0:v[0]]
		str2 := sql[v[1]-1 : len(sql)]
		sql = fmt.Sprintf("%s%s%s", str1, paramStr, str2)
		values[i] = paramGetter.Get(paramName)
	}

	return sql, values, nil
}
