package models

import (
	"fmt"
	"space-game/events"

	"github.com/google/uuid"
)

type GameObject interface {
	Update(...GameObjectUpdateOpt)
	ID() uuid.UUID
	RemoveHandler(uuid.UUID)
	HandlerRef(uuid.UUID, events.RemoveHandlerFunc)
}

type gameObject struct {
	tick       int64
	name       string
	id         uuid.UUID
	debugEvent *events.Event
	handlers   map[uuid.UUID][]events.RemoveHandlerFunc
}

type GameObjectOpts struct{}

var GameObjectOptions GameObjectOpts

type GameObjectOpt func(g *gameObject)

type GameObjectUpdateOpts struct{}

var GameObjectUpdateOptions GameObjectUpdateOpts

type GameObjectUpdateOpt func(g *gameObject)

func NewGameObject(opts ...GameObjectOpt) GameObject {
	g := gameObject{
		id:       uuid.New(),
		handlers: map[uuid.UUID][]events.RemoveHandlerFunc{},
	}
	for _, opt := range opts {
		opt(&g)
	}
	return g
}
func (g gameObject) Update(opts ...GameObjectUpdateOpt) {
	for _, opt := range opts {
		opt(&g)
	}

	if g.debugEvent != nil {
		g.debugEvent.Fire(fmt.Sprintf("%+v", g))
	}
}

func (g gameObject) ID() uuid.UUID {
	return g.id
}

func (g gameObject) HandlerRef(key uuid.UUID, rmv events.RemoveHandlerFunc) {
	if _, ok := g.handlers[key]; ok {
		g.handlers[key] = append(g.handlers[key], rmv)
		return
	}
	g.handlers[key] = []events.RemoveHandlerFunc{
		rmv,
	}
}

func (g gameObject) RemoveHandler(key uuid.UUID) {
	for _, rmv := range g.handlers[key] {
		if rmv != nil {
			rmv()
		}
	}
	g.handlers[key] = nil
}

// GameObjectOptions
func (GameObjectOpts) Name(n string) GameObjectOpt {
	return func(g *gameObject) {
		g.name = n
	}
}

func (GameObjectOpts) ID(id uuid.UUID) GameObjectOpt {
	return func(g *gameObject) {
		g.id = id
	}
}

func (GameObjectOpts) DebugEvent(e *events.Event) GameObjectOpt {
	return func(g *gameObject) {
		g.debugEvent = e
	}
}

// GameObjectUpdateOptions
func (GameObjectUpdateOpts) Tick(i int64) GameObjectUpdateOpt {
	return func(g *gameObject) {
		g.tick = i
	}
}
