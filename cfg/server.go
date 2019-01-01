package cfg

// Server holds Server config
type Server struct {
	Redis Redis

	// maximum number of running worker at a time
	ConcurrentWorker int
}
