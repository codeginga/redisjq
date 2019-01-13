package cnst

import "time"

const (
	// App holds app name
	App string = "redisjq"

	// Separator sperates/concates string
	Separator string = "#@#"

	// Lock is suffix of lock key
	Lock string = "_lock"

	// DefaultTaskRunTime represents default life time of running task
	DefaultTaskRunTime time.Duration = time.Second * 60

	// DefaultTaskMessageTTL represents default ttl of a message
	DefaultTaskMessageTTL time.Duration = time.Second * (60 * 10)

	// SleepDuration defines sleep time after run all worker
	SleepDuration time.Duration = time.Millisecond * 400

	// RedisEmptyMessage represents message for empty key
	// "github.com/go-redis/redis" returns this value for empty key
	RedisEmptyMessage string = "redis: nil"
)
