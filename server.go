package redisjq

import "context"

type server struct {
}

func (s *server) RegisterTask(name string, worker Worker) error {
	return nil
}

func (s *server) Start(ctx context.Context) error {
	return nil
}

// NewServer returns new instance of Server
func NewServer() Server {
	return &server{}
}
