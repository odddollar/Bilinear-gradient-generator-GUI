package widgets

import (
	"fmt"
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
	// Get widths and heights
	imgWidth := float32(p.DisplayImage.Bounds().Dx())
	imgHeight := float32(p.DisplayImage.Bounds().Dy())
	canvasWidth := p.renderer.image.Size().Width
	canvasHeight := p.renderer.image.Size().Height

	// Calculate aspect ratios
	imgAspect := imgWidth / imgHeight
	canvasAspect := canvasWidth / canvasHeight

	// Get new width and height of image
	var newWidth, newHeight float32
	if imgAspect > canvasAspect {
		// Image is wider relative to canvas
		newWidth = canvasWidth
		newHeight = canvasWidth / imgAspect
	} else {
		// Image is taller relative to canvas
		newWidth = canvasHeight * imgAspect
		newHeight = canvasHeight
	}

	// Calculate bounds
	top := (canvasHeight / 2) - (newHeight / 2)
	bottom := (canvasHeight / 2) + (newHeight / 2)
	left := (canvasWidth / 2) - (newWidth / 2)
	right := (canvasWidth / 2) + (newWidth / 2)

	fmt.Println("Clicked: ", event.Position)
	fmt.Println("Canvas size: ", canvasWidth, canvasHeight)
	fmt.Println("Canvas aspect: ", canvasAspect)
	fmt.Println("Image size: ", imgWidth, imgHeight)
	fmt.Println("Image aspect: ", imgAspect)
	fmt.Println("New image size: ", newWidth, newHeight)
	fmt.Printf("Top: %f, Bottom: %f, Left: %f, Right: %f\n", top, bottom, left, right)
	fmt.Println("---")

	clickedX := event.Position.X
	clickedY := event.Position.Y

	// Check point within image
	if clickedX >= left && clickedY >= top && clickedX < right && clickedY < bottom {
		// Get clicked position relative to image location
		relativeClickedX := clickedX - left
		relativeClickedY := clickedY - top

		// Map clicked relative position to values in original image
		mappedClickedX := mapRange(relativeClickedX, 0, newWidth, 0, imgWidth)
		mappedClickedY := mapRange(relativeClickedY, 0, newHeight, 0, imgHeight)

		fmt.Println(relativeClickedX, relativeClickedY, mappedClickedX, mappedClickedY)

		// Get NRGBA value a clicked point
		c := p.ActualImage.At(mappedClickedX, mappedClickedY).(color.NRGBA)
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
