package sqlExcuter

import (
	"database/sql"
	"torm/configs"
	"torm/context"
	"torm/parser"
	"torm/valueSetter"
)

func Query(config context.QueryConfig, c *context.DBQueryContext) (interface{}, error) {
	db, err := sql.Open(configs.GetDirver(), configs.GetConnection(config.DbKey))
	if err != nil {
		return nil, err
	}

	defer db.Close()
	sql, args, err := parser.Parser(config.Sql, c)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	columns, _ := rows.Columns()
	valueSetter := valueSetter.BuilderValueSetter(config, c)
	return valueSetter.Scan(config, c, rows, columns[:]), nil
}

func Update(config context.UpdateConfig, c *context.DBUpdateContext) error {
	db, err := sql.Open(configs.GetDirver(), configs.GetConnection(config.DbKey))
	if err != nil {
		return err
	}

	defer db.Close()
	sql, args, err := parser.Parser(config.Sql, c)
	if err != nil {
		return err
	}

	smt, err := db.Prepare(sql)
	if err != nil {
		return err
	}

	defer smt.Close()

	_, err = smt.Exec(args...)

	return err
}
