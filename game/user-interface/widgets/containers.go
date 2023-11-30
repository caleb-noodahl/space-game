package uiwidgets

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

func DisplayObjectContainer(obj interface{}, f font.Face) *widget.Container {
	bytes, err := json.Marshal(obj)
	if err != nil {
		log.Panic(err)
	}
	out := map[string]interface{}{}
	if err := json.Unmarshal(bytes, &out); err != nil {
		log.Panic(err)
	}
	sorted := []string{}
	for k := range out {
		sorted = append(sorted, k)
	}
	sort.Strings(sorted)

	container := NewDefaultObjectContainer()
	for _, s := range sorted {
		container.AddChild(NewDefaultLabel(s, f))
		container.AddChild(widget.NewLabel(
			widget.LabelOpts.TextOpts(widget.TextOpts.Insets(widget.NewInsetsSimple(5))),
			widget.LabelOpts.Text(fmt.Sprintf("%v", out[s]), f, &widget.LabelColor{
				Idle:     White,
				Disabled: DarkGrey,
			}),
		))
	}
	return container
}

func NewDefaultObjectContainer() *widget.Container {
	return widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Padding(widget.Insets{
				Left:  25,
				Right: 25,
			}),
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Stretch([]bool{false, true}, nil),
			widget.GridLayoutOpts.Spacing(20, 0),
		)))
}
