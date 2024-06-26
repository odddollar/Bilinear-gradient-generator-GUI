package widgets

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that extends image display with picker functionality
type ImageColourPicker struct {
	widget.BaseWidget
	renderer       *imageColourPickerRenderer
	tappedCallback func(color.NRGBA)
	cursor         desktop.Cursor

	// Holds internal image state
	actualImage  image.Image
	displayImage image.Image
	img          *canvas.Image
}

// Creates new ImageColourPicker widget.
// Callback is run when widget is clicked
func NewColourPicker(t func(color.NRGBA)) *ImageColourPicker {
	// Create empty canvas and set fill mode
	i := canvas.NewImageFromImage(nil)
	i.FillMode = canvas.ImageFillOriginal

	// Create new object with placeholder image and callback
	picker := &ImageColourPicker{
		cursor:         desktop.CrosshairCursor,
		tappedCallback: t,
		img:            i,
	}

	// Extend base widget and return
	picker.ExtendBaseWidget(picker)
	return picker
}

// Set widget's image. a is intended for the actual image, d for the image to be displayed.
// This allows an image to be displayed with a checkerboard background, whilst still using
// the actual image for picker functionality
func (p *ImageColourPicker) SetImage(a, d image.Image) {
	// Set internal images
	p.actualImage = a
	p.displayImage = d

	// Update displayed image and refresh
	p.img.Image = d
	p.img.Refresh()
}

// Get image colour at tapped point
func (p *ImageColourPicker) Tapped(event *fyne.PointEvent) {
	// Check if clicked inside actual image
	if inside, x, y := p.isInImage(event.Position); inside {
		// Get NRGBA value a clicked point
		c := p.actualImage.At(x, y).(color.NRGBA)
		p.tappedCallback(c)
	}
}

// Update cursor when it moves inside image
func (p *ImageColourPicker) MouseMoved(event *desktop.MouseEvent) {
	// Check if moved inside actual image
	if inside, _, _ := p.isInImage(event.Position); inside {
		p.cursor = desktop.CrosshairCursor
	} else {
		p.cursor = desktop.DefaultCursor
	}
}

// Does nothing but implements interface
func (p *ImageColourPicker) MouseIn(event *desktop.MouseEvent) {}

// Does nothing but implements interface
func (p *ImageColourPicker) MouseOut() {}

// Return current cursor
func (p *ImageColourPicker) Cursor() desktop.Cursor {
	return p.cursor
}

// Take position in canvas and check if in actual image, also returning mapped x and y
// within image
func (p *ImageColourPicker) isInImage(pos fyne.Position) (bool, int, int) {
	// Get widths and heights
	imgWidth := float32(p.displayImage.Bounds().Dx())
	imgHeight := float32(p.displayImage.Bounds().Dy())
	canvasWidth := p.renderer.img.Size().Width
	canvasHeight := p.renderer.img.Size().Height

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

	// Get position relative to image location
	relativeX := pos.X - left
	relativeY := pos.Y - top

	// Map relative position to values in original image
	mappedClickedX := mapRange(relativeX, 0, newWidth-1, 0, imgWidth-1)
	mappedClickedY := mapRange(relativeY, 0, newHeight-1, 0, imgHeight-1)

	// Check point within image
	if pos.X >= left && pos.Y >= top && pos.X < right && pos.Y < bottom {
		return true, mappedClickedX, mappedClickedY
	}
	return false, mappedClickedX, mappedClickedY
}

// Returns new renderer for ImageColourPicker
func (p *ImageColourPicker) CreateRenderer() fyne.WidgetRenderer {
	renderer := &imageColourPickerRenderer{
		icp: p,
		img: p.img,
	}
	p.renderer = renderer
	return renderer
}

// Renderer for ImageColourPicker widget
type imageColourPickerRenderer struct {
	icp *ImageColourPicker
	img *canvas.Image
}

// Returns minimum size of ImageColourPicker widget
func (r *imageColourPickerRenderer) MinSize() fyne.Size {
	return r.img.MinSize()
}

// Resizes image to fit available space
func (r *imageColourPickerRenderer) Layout(size fyne.Size) {
	r.img.Resize(size)
}

// Refreshes canvas on which image displayed
func (r *imageColourPickerRenderer) Refresh() {
	canvas.Refresh(r.icp)
}

// Returns child widgets of ImageColourPicker
func (r *imageColourPickerRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.img}
}

// Does nothing as ImageColourPicker doesn't hold any resources
func (r *imageColourPickerRenderer) Destroy() {}
