package backend

import "time"

// Locker handles locking
type Locker interface {
	Lock(key string) error
	Unlock(key string) error
}

// Set manages set operations
type Set interface {
	First() (key string, err error)
	Add(tim time.Time, key string) (err error)
	Remove(key string) error
}

// Task stores current status of the task
type Task interface {
	Save(key, value string) (err error)
	Get(key string) (value string, err error)
}

// Container contains all backend's interface
type Container struct {
	Locker Locker
	Set    Set
	Task   Task
}
