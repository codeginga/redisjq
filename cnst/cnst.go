package cnst

import "time"

const (
	// App holds app name
	App string = "redisjq"

	// DefaultTaskRunTime represents default life time of running task
	DefaultTaskRunTime time.Duration = time.Second * 60

	// DefaultTaskMessageTTL represents default ttl of a message
	DefaultTaskMessageTTL time.Duration = time.Second * (60 * 10)
)
