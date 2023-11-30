package main

import (
	"log"
	"space-game/game"
	"space-game/game/generator"
	userinterface "space-game/game/user-interface"
	"space-game/models"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// //
	// debug := events.NewEvent(
	// 	events.EventOptions.EventType("debug"),
	// 	events.EventOptions.Handlers([]events.EventHandler{*events.DefaultDebugEventHandler()}),
	// )

	player := models.NewPlayer(
		models.PlayerOptions.GameObject(
			models.NewGameObject(
				models.GameObjectOptions.Name("player"),
				//models.GameObjectOptions.DebugEvent(debug),
			),
		),
	)

	game := game.NewGame(
		game.GameOptions.EbitenUI(
			userinterface.NewEbitenUI(
				userinterface.EbitenUIOptions.ScreenSize(200, 400),
				userinterface.EbitenUIOptions.Title("space-game"),
			),
		),
		game.GameOptions.Generator(generator.NewGenerator(
			generator.GeneratorOptions.Stations(1),
			generator.GeneratorOptions.Player(player),
			//generator.GeneratorOptions.DebugEvent(debug),
		)),
	)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
