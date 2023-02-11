package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		data:      data,
		createdAt: now,
	}
}

func (e *Message) String() string {
	return fmt.Sprintf("Message %s", e.channel)
}

func (e *Message) Channel() Topic {
	return e.channel
}

func (e *Message) SetChannel(channel Topic) {
	e.channel = channel
}

func (e *Message) Data() interface{} {
	return e.data
}
