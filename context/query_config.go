package context

type Config struct {
	DbKey string
	Sql   string
	Type  int
}

type QueryConfig struct {
	Config
	Configs []QueryConfig
	Target  interface{}
}
