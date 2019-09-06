package context

import "database/sql"

type DBUpdateContext struct {
	DBContext
	UpdateConfig UpdateConfig
	txs          map[string]*sql.Tx
	openDbs      []*sql.DB
	openSmts     []*sql.Stmt
}

func (c *DBUpdateContext) GetTran(dbKey string) *sql.Tx {
	tx, ok := c.txs[dbKey]
	if ok {
		return tx
	}

	return nil
}

func (c *DBUpdateContext) AddTran(dbKey string, tx *sql.Tx) {
	if c.txs == nil {
		c.txs = make(map[string]*sql.Tx, 0)
	}

	c.txs[dbKey] = tx
}

func (c *DBUpdateContext) AddDB(db *sql.DB) {
	c.openDbs = append(c.openDbs, db)
}

func (c *DBUpdateContext) AddSmt(smt *sql.Stmt) {
	c.openSmts = append(c.openSmts, smt)
}

func (c *DBUpdateContext) RollBack() {
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

	c.txs = make(map[string]*sql.Tx, 0)
	c.openDbs = []*sql.DB{}
}

func (c *DBUpdateContext) Commit() {
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
