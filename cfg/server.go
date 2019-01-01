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
	RunTime int

	// Message TTL in seconds
	MessageTTL int

	// name of the task queue
	QName string
}

// RunTimeDuration converts RunTime in time.Duration
func (t *Task) RunTimeDuration() time.Duration {
	if t.RunTime == 0 {
		return cnst.DefaultTaskRunTime
	}

	return time.Second * time.Duration(t.RunTime)
}

// MessageTTLDuration converts MessageTTL in time.Duration
func (t *Task) MessageTTLDuration() time.Duration {
	if t.MessageTTL == 0 {
		return cnst.DefaultTaskMessageTTL
	}

	return time.Second * time.Duration(t.MessageTTL)
}
