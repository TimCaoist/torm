package context

import (
	"database/sql"
	"runtime"
)

type DBUpdateContext struct {
	DBContext
	UpdateConfig UpdateConfig
	txs          map[string]*sql.Tx
	openDbs      []*sql.DB
	openSmts     []*sql.Stmt
	relased      bool
}

func release(c *DBUpdateContext) {
	c.RollBack()
}

func NewDBUpdateContext() *DBUpdateContext {
	context := DBUpdateContext{}
	context.txs = make(map[string]*sql.Tx, 0)
	context.openDbs = make([]*sql.DB, 0)
	context.openSmts = make([]*sql.Stmt, 0)

	runtime.SetFinalizer(&context, release)
	return &context
}

func (c *DBUpdateContext) GetTran(dbKey string) *sql.Tx {
	tx, ok := c.txs[dbKey]
	if ok {
		return tx
	}

	return nil
}

func (c *DBUpdateContext) AddTran(dbKey string, tx *sql.Tx) {
	c.txs[dbKey] = tx
}

func (c *DBUpdateContext) AddDB(db *sql.DB) {
	c.openDbs = append(c.openDbs, db)
}

func (c *DBUpdateContext) AddSmt(smt *sql.Stmt) {
	c.openSmts = append(c.openSmts, smt)
}

func (c *DBUpdateContext) RollBack() {
	if c.relased == true {
		return
	}

	defer func() {
		c.closeDbs()
	}()

	for i, _ := range c.txs {
		c.txs[i].Rollback()
	}
}

func (c *DBUpdateContext) closeDbs() {
	for i, _ := range c.openDbs {
		c.openDbs[i].Close()
	}

	c.relased = true
	c.txs = nil
	c.openSmts = nil
	c.openDbs = nil
}

func (c *DBUpdateContext) Commit() {
	if c.relased == true {
		return
	}

	defer func() {
		c.closeDbs()
	}()

	for i, _ := range c.txs {
		c.txs[i].Commit()
	}

	for _, stmt := range c.openSmts {
		stmt.Close()
	}
}
