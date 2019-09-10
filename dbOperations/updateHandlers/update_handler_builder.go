package updateHandlers

import (
	"fmt"
	"torm/context"
)

const (
	Single_Insert = iota
	Single_Update
	Single_Delete
	Batch_Update
	Batch_Update_Filter
	Batch_Inert
	Batch_Delete
)

var updateHandlers map[int]IUpdateHandler

func Builder(config context.UpdateConfig) (*IUpdateHandler, error) {
	handler, ok := updateHandlers[config.Type]
	if ok {
		return &handler, nil
	}

	return nil, fmt.Errorf("Cann't found matching handler.")
}

func init() {
	updateHandlers = make(map[int]IUpdateHandler)
	updateHandlers[Single_Insert] = &SingleInsertHandler{}
	updateHandlers[Single_Update] = &SingleUpdateHandler{}
	updateHandlers[Single_Delete] = &SingleDeleteHandler{}
	updateHandlers[Batch_Update] = &BatchUpdateHandler{}
	updateHandlers[Batch_Update_Filter] = &BatchUpdateFilterHandler{}
	updateHandlers[Batch_Inert] = &BatchInsertHandler{}
	updateHandlers[Batch_Delete] = &BatchDeleteHandler{}
}
