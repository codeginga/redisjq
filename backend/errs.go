package backend

import "errors"

// list of backend errors
var (
	ErrExist = errors.New("exist")
	ErrEmpty = errors.New("empty")
)
