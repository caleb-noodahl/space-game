package models

import (
	"fmt"
	"math"
	"space-game/events"
	"space-game/game/utils"

	"github.com/google/uuid"
)

type StorageType int

const (
	Air StorageType = iota
	Water
	Food
)

func (s StorageType) String() string {
	switch s {
	case Air:
		return "air"
	case Water:
		return "water"
	case Food:
		return "food"
	}
	return ""
}

type Storage struct {
	Supply    float64     `json:"supply"`
	Capacity  float64     `json:"capacity"`
	Decrement float64     `json:"decrement"`
	Type      StorageType `json:"type"`
}

type EnvironmentalStorage struct {
	Owner        uuid.UUID `json:"owner"`
	Air          *Storage  `json:"air"`
	Food         *Storage  `json:"food"`
	Water        *Storage  `json:"water"`
	StorageEvent *events.Event
}

type StorageOpts struct{}

var StorageOptions StorageOpts

type StorageOpt func(s *Storage)

func NewStorage(opts ...StorageOpt) *Storage {
	s := &Storage{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s StorageOpts) Supply(sup float64) StorageOpt {
	return func(s *Storage) {
		s.Supply = sup
	}
}

func (s StorageOpts) Capacity(c float64) StorageOpt {
	return func(s *Storage) {
		s.Capacity = c
	}
}

func (s StorageOpts) Decrement(d float64) StorageOpt {
	return func(s *Storage) {
		s.Decrement = d
	}
}

func (s StorageOpts) Type(t StorageType) StorageOpt {
	return func(s *Storage) {
		s.Type = t
	}
}

type EnvironmentalStorageOpts struct{}

var EnvironmentalStorageOptions EnvironmentalStorageOpts

type EnvironmentalStorageOpt func(e *EnvironmentalStorage)

func NewEnvironmentalStorage(opts ...EnvironmentalStorageOpt) *EnvironmentalStorage {
	e := &EnvironmentalStorage{
		Air:   &Storage{},
		Food:  &Storage{},
		Water: &Storage{},
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e EnvironmentalStorageOpts) Owner(id uuid.UUID) EnvironmentalStorageOpt {
	return func(e *EnvironmentalStorage) {
		e.Owner = id
	}
}

func (e EnvironmentalStorageOpts) Air(a *Storage) EnvironmentalStorageOpt {
	return func(e *EnvironmentalStorage) {
		e.Air = a
		e.Air.Type = Air
	}
}

func (e EnvironmentalStorageOpts) Food(f *Storage) EnvironmentalStorageOpt {
	return func(e *EnvironmentalStorage) {
		e.Food = f
		e.Food.Type = Food
	}
}

func (e EnvironmentalStorageOpts) Water(w *Storage) EnvironmentalStorageOpt {
	return func(e *EnvironmentalStorage) {
		e.Water = w
		e.Water.Type = Water
	}
}

func (e EnvironmentalStorageOpts) Event(ev *events.Event) EnvironmentalStorageOpt {
	return func(e *EnvironmentalStorage) {
		e.StorageEvent = ev
	}
}

func (e *EnvironmentalStorage) Decrement(pop int) {
	e.Air.Supply = math.Round((e.Air.Supply-(e.Air.Decrement*float64(pop/3)))*100) / 100
	e.Food.Supply = math.Round((e.Food.Supply-(e.Food.Decrement*float64(pop/10)))*100) / 100
	e.Water.Supply = math.Round((e.Water.Supply-(e.Water.Decrement*float64(pop/2)))*100) / 100

	for _, storage := range []*Storage{e.Air, e.Food, e.Water} {
		if storage.Supply < 0 {
			storage.Supply = 0
		}
		switch {
		case storage.Supply <= 0:
			e.StorageEvent.FireOnce(events.NewEventArgs(
				events.EventArgsOptions.Level(events.Critical),
				events.EventArgsOptions.Hash(utils.Hash(fmt.Sprintf("%s%v%v", e.Owner.String(), events.Critical, storage.Type))),
				events.EventArgsOptions.Sender(e),
				events.EventArgsOptions.Payload(&events.EnvironmentalStorageNotificationPayload{
					StorageType: storage.Type.String(),
					Supply:      storage.Supply,
					Capacity:    storage.Capacity,
				}),
			))
		case storage.Supply < (storage.Capacity * 0.2):
			e.StorageEvent.FireOnce(events.NewEventArgs(
				events.EventArgsOptions.Level(events.Urgent),
				events.EventArgsOptions.Hash(utils.Hash(fmt.Sprintf("%s%v%v", e.Owner.String(), events.Urgent, storage.Type))),
				events.EventArgsOptions.Sender(e),
				events.EventArgsOptions.Payload(&events.EnvironmentalStorageNotificationPayload{
					StorageType: storage.Type.String(),
					Supply:      storage.Supply,
					Capacity:    storage.Capacity,
				}),
			))
		case storage.Supply < (storage.Capacity * 0.8):
			e.StorageEvent.FireOnce(events.NewEventArgs(
				events.EventArgsOptions.Hash(utils.Hash(fmt.Sprintf("%s%v%v", e.Owner.String(), events.Routine, storage.Type))),
				events.EventArgsOptions.Sender(e),
				events.EventArgsOptions.Level(events.Routine),
				events.EventArgsOptions.Payload(&events.EnvironmentalStorageNotificationPayload{
					StorageType: storage.Type.String(),
					Supply:      storage.Supply,
					Capacity:    storage.Capacity,
				}),
			))
		}
	}
}
