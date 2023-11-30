package generator

import (
	_ "embed"
	"fmt"
	"math/rand"

	"space-game/events"
	"space-game/game/utils"
	"space-game/models"
	npcs "space-game/models/npcs"

	"github.com/google/uuid"
)

type Generator struct {
	player      *models.Player
	debugEvent  *events.Event
	stationsCnt int
	gameObjects []models.GameObject
}

type GeneratorOpts struct{}

var GeneratorOptions GeneratorOpts

type GeneratorOpt func(g *Generator)

func NewGenerator(opts ...GeneratorOpt) *Generator {
	g := &Generator{
		gameObjects: []models.GameObject{},
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

func (g *Generator) GameObjects() []models.GameObject {
	//generate global objects
	//market := models.NewMarket()
	//generate stations
	for i := 0; i < g.stationsCnt; i++ {
		stationID := uuid.New()

		station := models.NewStation(
			models.StationOptions.HP(rand.Intn(5000-3000)+3000),
			models.StationOptions.Population(utils.Rand(12, 300)),
			models.StationOptions.GameObject(
				models.NewGameObject(
					models.GameObjectOptions.ID(stationID),
					models.GameObjectOptions.Name(fmt.Sprintf("station - %v", i)),
					models.GameObjectOptions.DebugEvent(g.debugEvent),
				),
			),
			models.StationOptions.EnvironmentalStorage(
				models.NewEnvironmentalStorage(
					models.EnvironmentalStorageOptions.Owner(stationID),
					models.EnvironmentalStorageOptions.Event(
						events.NewEvent(
							events.EventOptions.EventType("env:storage"),
						),
					),
					models.EnvironmentalStorageOptions.Air(
						models.NewStorage(
							models.StorageOptions.Capacity(1000),
							models.StorageOptions.Supply(1000),
							models.StorageOptions.Decrement(1),
						),
					),
					models.EnvironmentalStorageOptions.Water(
						models.NewStorage(
							models.StorageOptions.Capacity(1000),
							models.StorageOptions.Supply(1000),
							models.StorageOptions.Decrement(1),
						),
					),
					models.EnvironmentalStorageOptions.Food(
						models.NewStorage(
							models.StorageOptions.Capacity(200),
							models.StorageOptions.Supply(200),
							models.StorageOptions.Decrement(.2),
						),
					),
				),
			),
			models.StationOptions.StationNotificationEvent(events.NewEvent(
				events.EventOptions.EventType("st:notif"),
			)),
		)

		//generate personnel for each station
		manager := npcs.NewStationManager(
			npcs.StationManagerOptions.Station(station),
			npcs.StationManagerOptions.NPC(npcs.NewNPC(
				npcs.NPCOptions.GameObject(models.NewGameObject(
					models.GameObjectOptions.DebugEvent(g.debugEvent),
				)),
				npcs.NPCOptions.Character(
					models.NewCharacter(
						models.CharacterOptions.Background(models.Background(utils.Rand(0, 2))),
					),
				)),
			))

		//attach event handlers
		manager.HandlerRef(station.ID(), station.NotificationEvent.AddHandler(manager.HandleStationEvent))
		manager.HandlerRef(station.ID(), station.EnvironmentalStorage.StorageEvent.AddHandler(manager.HandleStationEvent))
		if i == 0 {
			g.player.HandleRef(station.ID(), station.NotificationEvent.AddHandler(g.player.HandleStationEvent))
		}
		g.gameObjects = append(g.gameObjects, station)
		g.gameObjects = append(g.gameObjects, manager)
	}
	//g.gameObjects = append(g.gameObjects, market)
	return g.gameObjects
}

func (g GeneratorOpts) Stations(cnt int) GeneratorOpt {
	return func(g *Generator) {
		g.stationsCnt = cnt
	}
}

func (g GeneratorOpts) DebugEvent(e *events.Event) GeneratorOpt {
	return func(g *Generator) {
		g.debugEvent = e
	}
}

func (g GeneratorOpts) Player(p *models.Player) GeneratorOpt {
	return func(g *Generator) {
		g.player = p
	}
}
