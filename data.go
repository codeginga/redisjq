package redisjq

// Message holds message for job queue
type Message struct {
	ID    string
	Name  string
	Retry int
	Value string

	delay int
}

// Delay retruns message pop delay
func (m *Message) Delay() int {
	return m.delay
}

// SetDelay sets message pop delay time
func (m *Message) SetDelay(sec int) {
	m.delay = sec
}

// Worker defines worker
type Worker func(Task)
