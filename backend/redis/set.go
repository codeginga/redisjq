package redis

import (
	"time"

	"github.com/codeginga/redisjq/backend"
)

type set struct {
}

func (s *set) First() (kye string, err error) {
	return
}

func (s *set) Add(tim time.Time, key string) (err error) {
	return
}

func (s *set) UpdateTim(tim time.Time, key string) (err error) {
	return
}

func (s *set) Remove(key string) (err error) {
	return
}

// NewSet returns instance of backend.Set
func NewSet() backend.Set {
	return &set{}
}
