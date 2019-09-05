package dataMapping

import (
	"reflect"
	"strings"
	"sync"
	"torm/common"
	_ "torm/common"
)

const (
	tormStr = "torm"
	split   = ";"
	split1  = ":"
	dbName  = "db_name"
	empty   = ""
)

var mapTypesOnce sync.Once

var mapTypes map[reflect.Type][]MappingData

func GetMapTypesInstance() map[reflect.Type][]MappingData {
	mapTypesOnce.Do(func() {
		mapTypes = make(map[reflect.Type][]MappingData)
	})

	return mapTypes
}

func GetTypeMapping(structType reflect.Type) []MappingData {
	realType := common.IndirectType(structType)
	instance := GetMapTypesInstance()
	caches, ok := instance[realType]
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

	instance[realType] = mappingDatas
	return mappingDatas
}

func GetMapingData(field reflect.StructField) MappingData {
	mappingData := MappingData{DBName: field.Name}
	tag := field.Tag
	var tormStr = tag.Get(tormStr)
	if tormStr == empty {
		return mappingData
	}

	strs := strings.Split(tormStr, split)
	for _, v := range strs {
		SetMapingField(v, mappingData, field)
	}

	return mappingData
}

func SetMapingField(config string, data MappingData, field reflect.StructField) {
	values := strings.Split(config, split1)
	switch values[0] {
	case dbName:
		data.DBName = values[1]
	}
}
