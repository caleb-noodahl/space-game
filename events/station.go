package events

import "github.com/google/uuid"

type StationNotificationEvent Event

type StationNotificationEventArgs struct {
	StationID uuid.UUID
	Level     EventLevel
	Msg       string
}
