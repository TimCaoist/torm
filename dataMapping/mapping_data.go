package dataMapping

type MappingData struct {
	DBName    string
	FieldName string
	Ingore    bool
}

var emptyMappingData MappingData

func GetMatchMapingData(dbName string, mapDatas []MappingData) (MappingData, bool) {
	for i, v := range mapDatas {
		if v.DBName == dbName {
			return mapDatas[i], true
		}
	}

	return emptyMappingData, false
}
