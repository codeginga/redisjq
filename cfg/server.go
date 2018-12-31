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
	// approximate time to run each task in seconds
	runTime *time.Duration

	// name of the task queue
	QName string
}

// SetRunTime sets approximate time of runing task in second
func (t *Task) SetRunTime(sec int) {
	d := time.Second * time.Duration(sec)
	t.runTime = &d
}

// RunTime returns time of running task
func (t *Task) RunTime() time.Duration {
	if t.runTime == nil {
		return cnst.DefaultTaskRunTime
	}

	return *t.runTime
}
