package context

type Config struct {
	DbKey    string
	Sql      string
	Type     int
	IsOnTran bool
}

type QueryConfig struct {
	Config
	Configs    []QueryConfig
	Target     interface{}
	OnlyOneRow bool
}
