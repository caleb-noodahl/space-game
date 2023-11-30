package game

import (
	"fmt"
	"space-game/events"
	"space-game/game/generator"
	userinterface "space-game/game/user-interface"
	"space-game/models"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ui           *userinterface.EbitenUI
	gen          *generator.Generator
	player       *models.Player
	init         int64
	current      int64
	clickedEvent *events.Event
	debugEvent   *events.Event
	destroyEvent *events.Event
	objects      []models.GameObject
}

type GameOpts struct{}

var GameOptions GameOpts

type GameOpt func(g *Game)

func NewGame(opts ...GameOpt) *Game {
	g := Game{
		init: time.Now().Unix(),
	}
	for _, opt := range opts {
		opt(&g)
	}
	return &g
}

// ebiten required function
func (g *Game) Update() error {
	//do higher level game loop stuff here
	g.current++
	for _, obj := range g.objects {
		obj.Update(
			models.GameObjectUpdateOptions.Tick(g.current),
		)
	}
	//update ui here
	if g.ui != nil {
		g.ui.Update()
	}

	//debug logic if the event is populated
	if g.debugEvent != nil {
		g.debugEvent.Fire(fmt.Sprintf("runtime: %v\nfps: %v",
			time.Now().Unix()-g.init,
			ebiten.ActualFPS(),
		))
	}
	return nil
}

// ebiten required function
func (g *Game) Draw(screen *ebiten.Image) {
	if g.ui != nil {
		g.ui.Draw(screen)
	}
}

// ebiten required function
func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g GameOpts) DebugEvent(e *events.Event) GameOpt {
	return func(g *Game) {
		g.debugEvent = e
	}
}

func (g GameOpts) DestroyEvent(e *events.Event) GameOpt {
	return func(g *Game) {
		g.destroyEvent = e
	}
}

func (g GameOpts) GameObjects(o []models.GameObject) GameOpt {
	return func(g *Game) {
		g.objects = o
	}
}

func (g GameOpts) Player(p *models.Player) GameOpt {
	return func(g *Game) {
		g.player = p
	}
}

func (g GameOpts) Generator(gen *generator.Generator) GameOpt {
	return func(g *Game) {
		g.gen = gen
		g.objects = g.gen.GameObjects()
	}
}

func (g GameOpts) EbitenUI(ui *userinterface.EbitenUI) GameOpt {
	return func(g *Game) {
		g.ui = ui
	}
}

func (g GameOpts) ClickedEvent(e *events.Event) GameOpt {
	return func(g *Game) {
		g.clickedEvent = e
	}
}
