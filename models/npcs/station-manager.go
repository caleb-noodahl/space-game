package models

import (
	"fmt"
	"log"
	"space-game/events"
	"space-game/models"

	"github.com/google/uuid"
)

type StationManager struct {
	npc     *NPC
	station *models.Station
}

// my self imposed penalty for not being able to find a creative solution
// to ideally i wish i could just do something like  - type StationManager NPC { station *Station }
func NewStationManager(opts ...StationManagerOpt) *StationManager {
	s := &StationManager{
		npc:     &NPC{},
		station: &models.Station{},
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

type StationManagerOpts struct{}

var StationManagerOptions StationManagerOpts

type StationManagerOpt func(s *StationManager)

func (StationManagerOpts) Station(station *models.Station) StationManagerOpt {
	return func(s *StationManager) {
		s.station = station
	}
}

func (StationManagerOpts) NPC(n *NPC) StationManagerOpt {
	return func(s *StationManager) {
		s.npc = n
	}
}

func (s StationManager) ID() uuid.UUID {
	return s.npc.gmObj.ID()
}

func (s *StationManager) Update(o ...models.GameObjectUpdateOpt) {
	s.npc.gmObj.Update(o...)
}

func (s StationManager) HandlerRef(key uuid.UUID, e events.RemoveHandlerFunc) {
	s.npc.gmObj.HandlerRef(key, e)
}

func (s StationManager) RemoveHandler(key uuid.UUID) {
	s.npc.gmObj.RemoveHandler(key)
}

func (s StationManager) HandleStationEvent(args interface{}) {
	e := args.(*events.EventArgs)
	switch e.Payload().(type) {
	//environmental
	case *events.EnvironmentalStorageNotificationPayload:
		payload := e.Payload().(*events.EnvironmentalStorageNotificationPayload)
		if e.Level() == events.Critical {
			s.station.NotificationEvent.Fire(
				events.NewEventArgs(
					events.EventArgsOptions.Level(events.Critical),
					events.EventArgsOptions.Payload(&events.StationNotificationEventArgs{
						StationID: s.station.ID(),
						Level:     events.Critical,
						Msg:       fmt.Sprintf("* ALERT * CRITICAL LEVELS OF %s", payload.StorageType),
					}),
				))
		}
	
		log.Printf(`station manager event handler - lvl : %v
		sender: %v
		sup: %v - cap: %v
		`, e.Level(), e.Sender(), payload.Supply, payload.Capacity)
	}
}
