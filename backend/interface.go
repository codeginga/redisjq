package backend

import "time"

// Locker handles locking
type Locker interface {
	Lock(key string) error
	Unlock(key string) error
}

// Set manages set operations
type Set interface {
	First() (kye string, err error)
	Add(tim time.Time, key string) (err error)
	UpdateTim(tim time.Time, key string) (err error)
	Remove(key string) error
}
