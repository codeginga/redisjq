package redisjq

import (
	"github.com/codeginga/redisjq/backend"
	"github.com/codeginga/redisjq/cnst"
)

type publisher struct {
	task  backend.Task
	set   backend.Set
	qname string
}

func (p *publisher) appQName() string {
	return cnst.App + "_" + p.qname
}

func (p *publisher) Publish(msg Message) (err error) {
	strMsg, err := msg.String()
	if err != nil {
		return
	}

	if err = p.task.Save(cnst.App+"_"+msg.ID, strMsg); err != nil {
		return
	}

	err = p.set.Add(msg.popupTime(), msg.ID)
	return
}

// NewPublisher returns new instance of Publisher
func NewPublisher() Publisher {
	return &publisher{}
}
