package context

type DBContext struct {
	Params      interface{}
	DBKey       string
	QueryConfig QueryConfig
}

type IDBContext interface {
	GetParams() interface{}
}

func (c DBContext) GetParams() interface{} {
	return c.Params
}
