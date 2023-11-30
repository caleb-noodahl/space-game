package models

import (
	"space-game/events"

	"github.com/google/uuid"
)

type Market struct {
	gameObj     GameObject
	MarketEvent *events.Event
	Contracts   []Contract
	Orders      []MarketOrder
}

func (m Market) ID() uuid.UUID {
	return m.gameObj.ID()
}

func (m Market) Update(opts ...GameObjectUpdateOpt) {
	m.gameObj.Update(opts...)
	//check the condition if any contracts
}

func (m Market) RemoveHandler(key uuid.UUID) {
	m.gameObj.RemoveHandler(key)
}

type MarketOpts struct {
}

var MarketOptions MarketOpts

type MarketOpt func(m *Market)

func NewMarket(opts ...MarketOpt) *Market {
	m := Market{}
	for _, opt := range opts {
		opt(&m)
	}
	return &m
}

func (m MarketOpts) GameObject(g GameObject) MarketOpt {
	return func(m *Market) {
		m.gameObj = g
	}
}
