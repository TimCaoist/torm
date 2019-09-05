package updateHandlers

import (
	"fmt"
	"torm/context"
)

var updateHandlers map[int]IUpdateHandler

func Builder(config context.UpdateConfig) (IUpdateHandler, error) {
	handler, ok := updateHandlers[config.Type]
	if ok {
		return handler, nil
	}

	return nil, fmt.Errorf("Can not found match handler.")
}

func init() {
	updateHandlers = make(map[int]IUpdateHandler)
	updateHandlers[1] = &SingleInsertHandler{}
}
