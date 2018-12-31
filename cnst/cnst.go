package cnst

import "time"

const (
	// App holds app name
	App string = "redisjq"

	// DefaultTaskLifeTime represents default life time of running task
	DefaultTaskLifeTime time.Duration = time.Second * 60
)
