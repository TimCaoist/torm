package updateHandlers

import "torm/context"

type SingleInsertHandler struct {
	UpdateHandler
}

func (qh SingleInsertHandler) Update(context *context.DBUpdateContext) error {

	return nil
}
