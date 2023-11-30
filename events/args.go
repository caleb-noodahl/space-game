package events

type EventArgs struct {
	level   EventLevel
	hash    string
	sender  interface{}
	payload interface{}
}

func (e *EventArgs) Hash() string {
	return e.hash
}

func (e *EventArgs) Payload() interface{} {
	return e.payload
}

func (e *EventArgs) Level() EventLevel {
	return e.level
}

func (e *EventArgs) Sender() interface{} {
	return e.sender
}

type EventArgsOpts struct{}

var EventArgsOptions EventArgsOpts

type EventArgsOpt func(e *EventArgs)

func NewEventArgs(opts ...EventArgsOpt) *EventArgs {
	e := &EventArgs{}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (EventArgsOpts) Level(lvl EventLevel) EventArgsOpt {
	return func(e *EventArgs) {
		e.level = lvl
	}
}

func (EventArgsOpts) Hash(hash string) EventArgsOpt {
	return func(e *EventArgs) {
		e.hash = hash
	}
}

func (EventArgsOpts) Sender(s any) EventArgsOpt {
	return func(e *EventArgs) {
		e.sender = s
	}
}
func (EventArgsOpts) Payload(p any) EventArgsOpt {
	return func(e *EventArgs) {
		e.payload = p
	}
}
