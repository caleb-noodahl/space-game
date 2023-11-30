package uiwidgets

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

type ListEntry struct {
	ID    string
	Title string
	Type  int
}

func NewDefaultList(entries []interface{}, f font.Face, e func(args *widget.ListEntrySelectedEventArgs)) *widget.List {
	return widget.NewList(
		//Set how wide the list should be
		widget.ListOpts.ContainerOpts(widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.MinSize(100, 0),
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				StretchVertical:    true,
			}),
		)),
		//Set the entries in the list
		widget.ListOpts.Entries(entries),
		widget.ListOpts.ScrollContainerOpts(
			//Set the background images/color for the list
			widget.ScrollContainerOpts.Image(&widget.ScrollContainerImage{
				Idle:     image.NewNineSliceColor(Background),
				Disabled: image.NewNineSliceColor(Background),
				Mask:     image.NewNineSliceColor(Background),
			}),
		),
		widget.ListOpts.SliderOpts(),
		widget.ListOpts.HideHorizontalSlider(),
		widget.ListOpts.HideVerticalSlider(),
		//Set the font for the list options
		widget.ListOpts.EntryFontFace(f),
		//Set the colors for the list
		widget.ListOpts.EntryColor(&widget.ListEntryColor{
			Selected:                   color.NRGBA{0, 255, 0, 100},     //Foreground color for the unfocused selected entry
			Unselected:                 White,                           //Foreground color for the unfocused unselected entry
			SelectedBackground:         color.NRGBA{100, 100, 100, 255}, //Background color for the unfocused selected entry
			SelectedFocusedBackground:  color.NRGBA{100, 100, 100, 255}, //Background color for the focused selected entry
			FocusedBackground:          color.NRGBA{100, 100, 100, 255}, //Background color for the focused unselected entry
			DisabledUnselected:         color.NRGBA{100, 100, 100, 255}, //Foreground color for the disabled unselected entry
			DisabledSelected:           color.NRGBA{100, 100, 100, 255}, //Foreground color for the disabled selected entry
			DisabledSelectedBackground: color.NRGBA{100, 100, 100, 255}, //Background color for the disabled selected entry
		}),
		//This required function returns the string displayed in the list
		widget.ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(ListEntry).Title
		}),
		//Padding for each entry
		widget.ListOpts.EntryTextPadding(widget.NewInsetsSimple(5)),
		//This handler defines what function to run when a list item is selected.
		widget.ListOpts.EntrySelectedHandler(e),
	)
}
