package userinterface

import (
	_ "embed"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed fonts/PxPlus_IBM_VGA_8x16.ttf
var ibm8x16 []byte

type EbitenUI struct {
	ui   *ebitenui.UI
	font font.Face
}

type EbitenUIOpts struct{}

var EbitenUIOptions EbitenUIOpts

type EbitenUIOpt func(e *EbitenUI)

func (e EbitenUI) Update() {

}

func (e EbitenUI) Draw(screen *ebiten.Image) {
	e.ui.Draw(screen)
}

func NewEbitenUI(opts ...EbitenUIOpt) *EbitenUI {
	tt, err := opentype.Parse(ibm8x16)
	if err != nil {
		log.Panic(err)
	}
	f, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	ui := &EbitenUI{
		font: f,
		ui: &ebitenui.UI{
			Container: DefaultRootContainer(),
		},
	}
	for _, opt := range opts {
		opt(ui)
	}

	return ui
}

func (e EbitenUIOpts) ScreenSize(height, width int) EbitenUIOpt {
	return func(e *EbitenUI) {
		ebiten.SetWindowSize(width, height)
	}
}

func (e EbitenUIOpts) Title(t string) EbitenUIOpt {
	return func(e *EbitenUI) {
		ebiten.SetWindowTitle(t)
	}
}
