package redis

import "github.com/codeginga/redisjq/backend"

type locker struct {
}

func (l *locker) Lock(key string) error {
	return nil
}

func (l *locker) Unlock(key string) error {
	return nil
}

// NewLocker returns instance of locker
func NewLocker() backend.Locker {
	return &locker{}
}
