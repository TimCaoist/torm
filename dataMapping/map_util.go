package dataMapping

import (
	"reflect"
	"strings"
	"torm/common"
	_ "torm/common"
)

const (
	tormStr = "torm"
	split   = ";"
	split1  = ":"
	dbName  = "db_name"
	ingore  = "ingore"
	empty   = ""
)

var mapTypes map[reflect.Type][]MappingData

func GetTypeMapping(structType reflect.Type) []MappingData {
	realType := common.IndirectType(structType)
	caches, ok := mapTypes[realType]
	if ok {
		return caches
	}

	filedCount := realType.NumField()
	mappingDatas := make([]MappingData, 0)
	for i := 0; i < filedCount; i++ {
		field := realType.Field(i)
		if field.Anonymous {
			continue
		}

		mappingDatas = append(mappingDatas, GetMapingData(field))
	}

	mapTypes[realType] = mappingDatas
	return mappingDatas
}

func GetMapingData(field reflect.StructField) MappingData {
	mappingData := &MappingData{DBName: field.Name, FieldName: field.Name}
	tag := field.Tag
	var tormStr = tag.Get(tormStr)
	if tormStr == empty {
		return *mappingData
	}

	strs := strings.Split(tormStr, split)
	for _, v := range strs {
		SetMapingField(v, mappingData, field)
	}

	return *mappingData
}

func SetMapingField(config string, data *MappingData, field reflect.StructField) {
	values := strings.Split(config, split1)
	switch values[0] {
	case dbName:
		data.DBName = values[1]
	case ingore:
		data.Ingore = true
	}
}

func init() {
	mapTypes = map[reflect.Type][]MappingData{}
}
