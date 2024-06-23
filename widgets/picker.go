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
	// Image will be manually assigned
	Image    image.Image
	renderer *imageColourPickerRenderer
}

// Creates new ImageColourPicker widget
func NewColourPicker() *ImageColourPicker {
	picker := &ImageColourPicker{}
	picker.ExtendBaseWidget(picker)
	return picker
}

// Returns new rendered for ImageColourPicker
func (p *ImageColourPicker) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromImage(p.Image)
	img.FillMode = canvas.ImageFillOriginal
	renderer := &imageColourPickerRenderer{
		image:   img,
		picker:  p,
		objects: []fyne.CanvasObject{img},
	}
	p.renderer = renderer
	return renderer
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

// Refresh the widget and its children/renderer
func (p *ImageColourPicker) Refresh() {
	p.BaseWidget.Refresh()
	p.renderer.Refresh()
}

// Renderer for ImageColourPicker widget
type imageColourPickerRenderer struct {
	picker  *ImageColourPicker
	image   *canvas.Image
	objects []fyne.CanvasObject
}

// Returns minimum size of ImageColourPicker widget
func (r *imageColourPickerRenderer) MinSize() fyne.Size {
	return r.image.MinSize()
}

// Resizes image to fit available space
func (r *imageColourPickerRenderer) Layout(size fyne.Size) {
	r.image.Resize(size)
}

// Refreshes canvas on which image displayed
func (r *imageColourPickerRenderer) Refresh() {
	// Set displayed canvas's image to image of picker widget
	r.image.Image = r.picker.Image
	canvas.Refresh(r.image)
}

// Returns child widgets of ImageColourPicker
func (r *imageColourPickerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// Does nothing as ImageColourPicker doesn't hold any resources
func (r *imageColourPickerRenderer) Destroy() {}
