package global

import (
	"Bilinear-gradient-generator-GUI/widgets"
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// Holds global image state
var ImageCurrent image.Image

// Holds corner pixels
var (
	TopLeftPixel     color.NRGBA
	TopRightPixel    color.NRGBA
	BottomLeftPixel  color.NRGBA
	BottomRightPixel color.NRGBA
)

// Variables to hold widgets
var (
	A                 fyne.App
	MainWindow        fyne.Window
	ImageDisplay      *widgets.ImageColourPicker
	RandomiseButton   *widget.Button
	SaveButton        *widget.Button
	OptionsButton     *widget.Button
	AboutButton       *widget.Button
	TopLeftButton     *widget.Button
	TopRightButton    *widget.Button
	BottomLeftButton  *widget.Button
	BottomRightButton *widget.Button
	SpacerWidget      *widgets.Spacer
)

// Variables to hold different layouts
// One holds layout with corner buttons, the other without
var (
	CornerButtonContent   *fyne.Container
	NoCornerButtonContent *fyne.Container
)
