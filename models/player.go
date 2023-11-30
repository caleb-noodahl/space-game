package models

import (
	"log"
	"space-game/events"

	"github.com/google/uuid"
)

type Player struct {
	character *Character
	gameObj   GameObject
}

type PlayerOpts struct{}

var PlayerOptions PlayerOpts

type PlayerOpt func(p *Player)

func NewPlayer(opts ...PlayerOpt) *Player {
	p := &Player{}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p Player) ID() uuid.UUID {
	return p.gameObj.ID()
}

func (p Player) Update(o ...GameObjectUpdateOpt) {
	p.gameObj.Update(o...)
}

func (p Player) HandleRef(key uuid.UUID, e events.RemoveHandlerFunc) {
	p.gameObj.HandlerRef(key, e)
}

func (p Player) HandleStationEvent(args interface{}) {
	e := args.(*events.EventArgs)
	switch e.Payload().(type) {
	case *events.StationNotificationEventArgs:
		event := e.Payload().(*events.StationNotificationEventArgs)
		if event.Level == events.Critical {
			log.Printf("player responding to %s event: level %v - %s", event.StationID, event.Level, event.Msg)
		}
	}
}

func (p PlayerOpts) GameObject(g GameObject) PlayerOpt {
	return func(p *Player) {
		p.gameObj = g
	}
}

func (p PlayerOpts) Character(c *Character) PlayerOpt {
	return func(p *Player) {
		p.character = c
	}
}
