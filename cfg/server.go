package cfg

import (
	"time"

	"github.com/codeginga/redisjq/cnst"
)

// Server holds Server config
type Server struct {
	Redis Redis

	// maximum number of running worker at a time
	ConcurrentWorker int
}

// Task holds task config
type Task struct {
	// approximate to run each task in seconds
	lifeTime *time.Duration
}

// SetLifeTime sets task aproximate life time in seconds
func (t *Task) SetLifeTime(sec int) {
	d := time.Second * time.Duration(sec)
	t.lifeTime = &d
}

// LifeTime returns task's life time in second
func (t *Task) LifeTime() time.Duration {
	if t.lifeTime == nil {
		return cnst.DefaultTaskLifeTime
	}

	return *t.lifeTime
}
