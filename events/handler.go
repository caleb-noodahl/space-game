package events

import "github.com/google/uuid"

type HandlerFunc func(args interface{})

type EventHandler struct {
	index int
	id    uuid.UUID
	h     HandlerFunc
	t     string
}

type EventHandlerOpts struct{}

var EventHandlerOptions EventHandlerOpts

type EventHandlerOpt func(e *EventHandler)

func NewEventHandler(opts ...EventHandlerOpt) *EventHandler {
	e := &EventHandler{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (EventHandlerOpts) Type(t string) EventHandlerOpt {
	return func(e *EventHandler) {
		e.t = t
	}
}

func (EventHandlerOpts) ID(id uuid.UUID) EventHandlerOpt {
	return func(e *EventHandler) {
		e.id = id
	}
}

func (EventHandlerOpts) HandlerFunc(h HandlerFunc) EventHandlerOpt {
	return func(e *EventHandler) {
		e.h = h
	}
}
