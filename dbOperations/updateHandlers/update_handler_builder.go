package updateHandlers

import (
	"fmt"
	"torm/context"
)

const (
	Single_Insert       = 1
	Single_Update       = 2
	Batch_Update        = 3
	Batch_Update_Filter = 4
	Batch_Inert         = 5
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
	updateHandlers[Batch_Update] = &BatchUpdateHandler{}
	updateHandlers[Batch_Update_Filter] = &BatchUpdateFilterHandler{}
	updateHandlers[Batch_Inert] = &BatchInsertHandler{}
}
