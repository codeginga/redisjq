package redisjq

import (
	"encoding/json"
	"time"

	"github.com/codeginga/redisjq/cnst"
)

// Message holds message for job queue
type Message struct {
	// set qniqueue id of each task
	ID string `json:"id"`

	// Name defines task name
	// set uniqueue name for different kind of  task
	Name string `json:"name"`

	Retry int    `json:"retry"`
	Value string `json:"value"`

	// task pop delay time in seconds
	Delay int `json:"delay"`
}

func (m *Message) popupTime() time.Time {
	t := time.Now().Add(time.Second * time.Duration(m.Delay))
	return t
}

// Key generates key for message
func (m *Message) Key() string {
	return cnst.App + cnst.Separator + m.ID
}

// String converts Message to string
func (m *Message) String() (val string, err error) {
	bts, err := json.Marshal(m)
	if err != nil {
		return
	}

	val = string(bts)
	return
}

// FrmString converts string to Message
func (m *Message) FrmString(str string) (err error) {
	err = json.Unmarshal([]byte(str), m)
	return
}

// Worker defines worker
type Worker func(Task)
