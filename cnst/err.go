package cnst

import "errors"

var (
	// ErrLocked , happens when key is locked by another one
	ErrLocked = errors.New("locked")
	// ErrEmptyQ , happens when job queue is empty
	ErrEmptyQ = errors.New("empty_queue")
	// ErrEmptyTask , happens when task value is empty
	ErrEmptyTask = errors.New("empty_task")
	// ErrRetryExit , happens when retry count is less than 0
	ErrRetryExit = errors.New("retry_exit")
)
