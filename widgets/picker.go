package widgets

import (
	"image"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that extends image display with picker functionality
type ImageColourPicker struct {
	widget.BaseWidget
	Image image.Image
}

// Creates new ImageColourPicker widget
func NewColourPicker(img image.Image) *ImageColourPicker {
	picker := &ImageColourPicker{Image: img}
	picker.ExtendBaseWidget(picker)
	return picker
}

// Returns new rendered for ImageColourPicker
func (p *ImageColourPicker) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromImage(p.Image)
	img.FillMode = canvas.ImageFillOriginal
	return widget.NewSimpleRenderer(img)
}

// Get image colour at tapped point
func (p *ImageColourPicker) Tapped(event *fyne.PointEvent) {
	// Get cursor click position
	x := int(event.Position.X)
	y := int(event.Position.Y)

	// Check point within image
	if x >= 0 && y >= 0 && x < p.Image.Bounds().Dx() && y < p.Image.Bounds().Dy() {
		// Get NRGBA value a clicked point
		col := p.Image.At(x, y).(color.NRGBA)
		log.Printf("Clicked at (%d, %d), Colour: %#v\n", x, y, col)
	}
}
