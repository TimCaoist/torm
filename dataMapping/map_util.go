package dataMapping

import "reflect"

func GetTypeMapping(structType reflect.Type) []MappingData {
	filedCount := structType.NumField()
	mappingDatas := make([]MappingData, 0)
	for i := 0; i < filedCount; i++ {
		field := structType.Field(i)
		if field.Anonymous {
			continue
		}

		//tag := field.Tag
		mappingData := &MappingData{DBName: field.Name}
		mappingDatas = append(mappingDatas, *mappingData)
	}

	return mappingDatas
}

func GetFieldName(tag reflect.StructTag, field reflect.StructField) (string, string) {
	// if tag.Get("s") {
	// 	return field.Name[:], field.Name[:]
	// }

	return "", ""
}
