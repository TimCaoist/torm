package updateHandlers

import (
	"torm/context"
)

type UpdateHandler struct {
	Excuter IUpdateHandler
}

type IUpdateHandler interface {
	Update(context *context.DBUpdateContext) error
}

func (qh UpdateHandler) Update(context *context.DBUpdateContext) error {
	return qh.Excuter.Update(context)
}

func GetUpdateHandler(context *context.DBUpdateContext) UpdateHandler {
	q := UpdateHandler{}
	handler, err := Builder(context.UpdateConfig)
	if err != nil {
		panic(err)
	}

	q.Excuter = handler
	return q
}
