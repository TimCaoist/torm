package parser

import (
	"errors"
	"regexp"
	"torm/context"
	"torm/paramUtil"
)

const (
	paramReg  = "@.*?[, ]|@.*?[)]"
	whieSpace = " "
)

func Parser(sql string, context context.IDBContext) (string, []interface{}, error) {
	reg := regexp.MustCompile(paramReg)
	//返回匹配到的结果
	result := reg.FindAllStringIndex(sql, -1)
	paramGetter := paramUtil.GetParamGetter(context)
	if paramGetter == nil {
		err := errors.New("Not Support")
		return sql, nil, err
	}

	sql, args := paramGetter.GetArgs(result, sql)
	return sql, args, nil
}
