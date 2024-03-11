package stack

import (
	"github.com/nats-io/nats.go/jetstream"
	"github.com/udonetsm/investing/models"
)

type Stack models.Stack

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(m *models.Messages) {
	s.Mess = append(s.Mess[:s.Count], m)
	s.Count++
}

func (s *Stack) Pop() *models.Messages {
	if s.Count == 0 {
		return &models.Messages{Error: "empty stack"}
	}
	s.Count--
	return s.Mess[s.Count]
}

func (s *Stack) Add(msg any) {
	if jmsg, ok := msg.(jetstream.Msg); ok {
		s.addAsJetsreamMsg(jmsg)
	}
	// if use some other broker, write new method for it.
	// Example(package rabbitmq is not real package. Use your real target returning type of message):
	// if rmqms, ok := msg.(rabbitmq.Msg); ok {
	// 	s.addAsRebbitMQ(rmqmsg)
	// }
}

func (s *Stack) addAsJetsreamMsg(msg jetstream.Msg) {
	m := models.Messages{}
	meta, metaerr := msg.Metadata()
	if metaerr != nil {
		m.Error = metaerr.Error()
		s.Push(&models.Messages{Error: metaerr.Error()})
	}
	m.Message = string(msg.Data())
	m.Timestamp = meta.Timestamp.UnixNano()
	s.Push(&models.Messages{Message: m.Message, Timestamp: m.Timestamp})
}
