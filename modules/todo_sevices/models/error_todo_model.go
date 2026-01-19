package todoModels

import "errors"

var (
	ErrNotFoundTodo   = errors.New("error: todo not found.")
	ErrParseUUIDError = errors.New("error: uuid invalid")
)
