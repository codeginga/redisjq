package cnst

// Status defines task status
type Status string

func (s Status) String() string {
	return string(s)
}

// list possible task status
const (
	RunningStatus Status = "running"
	DoneStatus    Status = "done"
	PendingStatus Status = "pending"
)
