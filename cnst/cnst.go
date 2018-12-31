package cnst

import "time"

const (
	// App holds app name
	App string = "redisjq"

	// DefaultTaskRunTime represents default life time of running task
	DefaultTaskRunTime time.Duration = time.Second * 60
)
