package backend

import "errors"

// list of backend errors
var (
	ErrLocked = errors.New("locked")
	ErrEmpty  = errors.New("empty")
)
