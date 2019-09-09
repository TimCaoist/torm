package sqlExcuter

import (
	"database/sql"
	"torm/configs"
	"torm/context"
	"torm/parser"
	"torm/valueSetter"
)

func Query(config *context.QueryConfig, c *context.DBQueryContext) (interface{}, error) {
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
	if config.OnlyOneRow == false {
		return valueSetter.Scan(config, c, rows, columns[:]), nil
	}

	return valueSetter.ScanOneRow(config, c, rows, columns[:]), nil
}

func Update(config *context.UpdateConfig, c *context.DBUpdateContext) error {
	if config.IsOnTran {
		return UpdateOnTran(config, c)
	}

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

	res, err := smt.Exec(args...)
	if config.RequireId == false {
		return err
	}

	id, err := res.LastInsertId()
	config.Id = id
	return err
}

func UpdateOnTran(config *context.UpdateConfig, c *context.DBUpdateContext) error {
	dbKey := config.DbKey
	tx := c.GetTran(dbKey)
	if tx == nil {
		db, err := sql.Open(configs.GetDirver(), configs.GetConnection(config.DbKey))
		if err != nil {
			return err
		}

		tx1, err := db.Begin()
		if err != nil {
			return err
		}

		c.AddTran(config.DbKey, tx1)
		c.AddDB(db)
		tx = tx1
	}

	sql, args, err := parser.Parser(config.Sql, c)
	if err != nil {
		c.RollBack()
		return err
	}

	smt, err := tx.Prepare(sql)
	if err != nil {
		c.RollBack()
		return err
	}

	c.AddSmt(smt)
	res, err := smt.Exec(args...)
	if err != nil {
		c.RollBack()
		return err
	}

	if config.RequireId == false {
		return err
	}

	id, err := res.LastInsertId()
	config.Id = id
	return err
}
