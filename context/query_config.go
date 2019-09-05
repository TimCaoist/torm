package context

type QueryConfig struct {
	DbKey   string
	Sql     string
	Configs []QueryConfig
	Target  interface{}
}
