package configs

type DBConfig struct {
	Connections map[string]string
	Driver      string
}

var dbConfig DBConfig

func init() {
	dbConfig = DBConfig{}
	dbConfig.Connections = make(map[string]string)
}

func UseDriver(driver string) {
	dbConfig.Driver = driver
}

func GetDirver() string {
	return dbConfig.Driver
}

func RegisterConnection(name, connectionStr string) {
	dbConfig.Connections[name] = connectionStr
}

func GetConnection(source string) string {
	val, ok := dbConfig.Connections[source]
	if ok {
		return val
	}

	return source
}
