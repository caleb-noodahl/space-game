package models

import "github.com/google/uuid"

type Equipment struct {
	ID      uuid.UUID `json:"uuid"`
	OwnerID uuid.UUID `json:"owner_id"`
	Name    string    `json:"name"`
}
