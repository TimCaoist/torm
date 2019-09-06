package context

import "reflect"

type UpdateConfig struct {
	Config
	UpdateModel UpdateModel
	RequireId   bool
	Id          int64
	KeyField    reflect.Value
}
