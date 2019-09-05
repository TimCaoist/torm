package common

import "reflect"

func Indirect(reflectValue reflect.Value) reflect.Value {
	for reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	return reflectValue
}

func IndirectType(reflectType reflect.Type) reflect.Type {
	for reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	return reflectType
}
