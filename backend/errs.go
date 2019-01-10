package backend

import "errors"

// list of backend errors
var (
	ErrLocked    = errors.New("locked")
	ErrEmptyQ    = errors.New("empty_queue")
	ErrEmptyTask = errors.New("empty_task")
)
