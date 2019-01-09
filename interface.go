package redisjq

import "context"

// Publisher wraps message publish
type Publisher interface {
	Publish(msg Message) error
}

// Server wraps server methods
type Server interface {
	RegisterTask(name string, worker Worker) error
	Start(ctx context.Context) error
}

// Task wraps task methods
type Task interface {
	Message() Message
	Done() error
	Retry() error
}
