package events

import (
	"log"

	"github.com/google/uuid"
)

var DebugEvent *Event

// meant to provide the simplest examples of an event handler
func DefaultDebugEventHandler() *EventHandler {
	return NewEventHandler(
		EventHandlerOptions.ID(uuid.Nil),
		EventHandlerOptions.HandlerFunc(func(args interface{}) {
			log.Printf("%v", args)
		}),
	)
}
