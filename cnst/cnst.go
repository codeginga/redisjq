package cnst

import "time"

const (
	// App holds app name
	App string = "redisjq"

	// Separator sperates/concates string
	Separator string = "#@#"

	// DefaultTaskRunTime represents default life time of running task
	DefaultTaskRunTime time.Duration = time.Second * 60

	// DefaultTaskMessageTTL represents default ttl of a message
	DefaultTaskMessageTTL time.Duration = time.Second * (60 * 10)

	// RedisEmptyMessage represents message for empty key
	// "github.com/go-redis/redis" returns this value for empty key
	RedisEmptyMessage string = "redis: nil"
)
