package redisjq

type task struct {
}

func (t *task) Message() Message {
	return Message{}
}

func (t *task) Done() error {
	return nil
}

func (t *task) Retry() error {
	return nil
}

func (t *task) RetryDelay() error {
	return nil
}
