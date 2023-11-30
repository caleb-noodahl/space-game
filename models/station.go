package models

import (
	"log"
	"space-game/events"

	"github.com/google/uuid"
)

type Station struct {
	gmObj                GameObject
	NotificationEvent    *events.Event
	EnvironmentalStorage *EnvironmentalStorage
	Market               *Market
	Name                 string `json:"name"`
	HP                   int    `json:"hp"`
	Population           int    `json:"pop"`
}

type StationOpts struct{}

var StationOptions StationOpts

type StationOpt func(s *Station)

func NewStation(opts ...StationOpt) *Station {
	o := Station{
		HP: 3000,
		EnvironmentalStorage: &EnvironmentalStorage{
			Air:   &Storage{},
			Food:  &Storage{},
			Water: &Storage{},
		},
	}
	for _, opt := range opts {
		opt(&o)
	}
	return &o
}

func (s Station) ID() uuid.UUID {
	return s.gmObj.ID()
}

func (s Station) RemoveHandler(key uuid.UUID) {
	s.gmObj.RemoveHandler(key)
}

func (s Station) HandlerRef(key uuid.UUID, e events.RemoveHandlerFunc) {
	s.gmObj.HandlerRef(key, e)
}

func (s *Station) Update(o ...GameObjectUpdateOpt) {
	s.gmObj.Update(o...)
	//simulate population using the station's resources
	s.EnvironmentalStorage.Decrement(s.Population)
	//debug todo remove
	log.Printf("%s : food: %v, water: %v, air: %v",
		s.ID(),
		s.EnvironmentalStorage.Food.Supply,
		s.EnvironmentalStorage.Water.Supply,
		s.EnvironmentalStorage.Air.Supply,
	)
}

func (StationOpts) GameObject(g GameObject) StationOpt {
	return func(s *Station) {
		s.gmObj = g
	}
}

func (StationOpts) StationNotificationEvent(e *events.Event) StationOpt {
	return func(s *Station) {
		s.NotificationEvent = e
	}
}

func (StationOpts) HP(hp int) StationOpt {
	return func(s *Station) {
		s.HP = hp
	}
}

func (StationOpts) Population(p int) StationOpt {
	return func(s *Station) {
		s.Population = p
	}
}

func (StationOpts) EnvironmentalStorage(e *EnvironmentalStorage) StationOpt {
	return func(s *Station) {
		s.EnvironmentalStorage = e
	}
}

func (StationOpts) Market(m *Market) StationOpt {
	return func(s *Station) {
		s.Market = m
	}
}
