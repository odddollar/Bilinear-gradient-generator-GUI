package widgets

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// TODO: Make widget resize properly

// Custom widget that extends image display with picker functionality
type ImageColourPicker struct {
	widget.BaseWidget
	renderer       *imageColourPickerRenderer
	tappedCallback func(color.NRGBA)

	// Images will be manually assigned
	ActualImage  image.Image
	DisplayImage image.Image
}

// Creates new ImageColourPicker widget.
// Callback is run when widget is clicked
func NewColourPicker(t func(color.NRGBA)) *ImageColourPicker {
	picker := &ImageColourPicker{tappedCallback: t}
	picker.ExtendBaseWidget(picker)
	return picker
}

// Returns new renderer for ImageColourPicker
func (p *ImageColourPicker) CreateRenderer() fyne.WidgetRenderer {
	img := canvas.NewImageFromImage(p.DisplayImage)
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
	if x >= 0 && y >= 0 && x < p.ActualImage.Bounds().Dx() && y < p.ActualImage.Bounds().Dy() {
		// Get NRGBA value a clicked point
		c := p.ActualImage.At(x, y).(color.NRGBA)
		p.tappedCallback(c)
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
	r.image.Image = r.picker.DisplayImage
	canvas.Refresh(r.image)
}

// Returns child widgets of ImageColourPicker
func (r *imageColourPickerRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

// Does nothing as ImageColourPicker doesn't hold any resources
func (r *imageColourPickerRenderer) Destroy() {}
