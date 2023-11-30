package events

import (
	"slices"

	"github.com/google/uuid"
)

type EventLevel int

const (
	Critical EventLevel = iota
	Urgent
	Routine
	Notification
)

type Event struct {
	hashes    []string
	eventType string
	handlers  []EventHandler
}

type EventOpts struct{}

var EventOptions EventOpts

type EventOpt func(e *Event)

type RemoveHandlerFunc func()

func NewEvent(opts ...EventOpt) *Event {
	e := &Event{
		hashes: []string{},
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e EventOpts) EventType(t string) EventOpt {
	return func(e *Event) {
		e.eventType = t
	}
}

func (e EventOpts) Handlers(handlers []EventHandler) EventOpt {
	return func(e *Event) {
		e.handlers = handlers
	}
}

func (e *Event) RemoveHandler(id uuid.UUID) {
	index := 0
	for i, handler := range e.handlers {
		if handler.id == id {
			index = i
			break
		}
	}
	e.handlers = append(e.handlers[:index], e.handlers[index+1:]...)
}

func (e *Event) AddHandler(h HandlerFunc) RemoveHandlerFunc {
	id := uuid.New()
	e.handlers = append(e.handlers, EventHandler{
		id: id,
		h:  h,
	})
	return func() { e.RemoveHandler(id) }
}

func (e *Event) Fire(args interface{}) {
	for _, h := range e.handlers {
		h.h(args)
	}
}

func (e *Event) FireOnce(args interface{}) {
	a := args.(*EventArgs)
	if slices.Contains(e.hashes, a.hash) {
		return
	}
	e.Fire(args)
	e.hashes = append(e.hashes, a.hash)
}
