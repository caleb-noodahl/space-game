package uiwidgets

import (
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

func NewDefaultLabel(text string, f font.Face) *widget.Label {
	return widget.NewLabel(
		widget.LabelOpts.TextOpts(widget.TextOpts.Insets(widget.NewInsetsSimple(5))),
		widget.LabelOpts.Text(text, f, &widget.LabelColor{
			Idle:     DarkGrey,
			Disabled: DarkGrey,
		}),
	)
}
