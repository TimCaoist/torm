package parser

import (
	"regexp"
	"torm/context"
	"torm/paramUtil"
)

const (
	paramReg = "@.*?[, ]|@.*?[);]"
)

func Parser(sql string, context context.IDBContext) (string, []interface{}, error) {
	reg := regexp.MustCompile(paramReg)
	//返回匹配到的结果
	result := reg.FindAllStringIndex(sql, -1)
	paramGetter := paramUtil.GetParamGetter(context)
	if paramGetter == nil {
		args := make([]interface{}, 0)
		return sql, args, nil
	}

	sql, args := paramGetter.GetArgs(result, sql)
	return sql, args, nil
}
