package redisjq

type publisher struct {
}

func (p *publisher) Publish(msg Message) error {
	return nil
}

// NewPublisher returns new instance of Publisher
func NewPublisher() Publisher {
	return &publisher{}
}
