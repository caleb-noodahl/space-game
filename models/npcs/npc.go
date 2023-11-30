package models

import (
	"space-game/models"

	"github.com/google/uuid"
)

type NPC struct {
	Character *models.Character
	Location  string `json:"location"`
	gmObj     models.GameObject
	listeners map[string][]func()
}

type NPCOpts struct{}

var NPCOptions NPCOpts

type NPCOpt func(n *NPC)

func NewNPC(opts ...NPCOpt) *NPC {
	n := &NPC{
		gmObj: models.NewGameObject(),
	}
	for _, opt := range opts {
		opt(n)
	}
	return n
}

func (NPCOpts) Character(c *models.Character) NPCOpt {
	return func(n *NPC) {
		n.Character = c
	}
}

func (NPCOpts) GameObject(g models.GameObject) NPCOpt {
	return func(n *NPC) {
		n.gmObj = g
	}
}

func (n *NPC) Update(o ...models.GameObjectUpdateOpt) {
	n.gmObj.Update(o...)
}

func (n NPC) ID() uuid.UUID {
	return n.gmObj.ID()
}
