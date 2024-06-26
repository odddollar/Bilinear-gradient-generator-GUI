package widgets

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that extends image display with picker functionality
type ImageColourPicker struct {
	widget.BaseWidget
	//renderer *imageColourPickerRenderer
	tappedCallback func(color.NRGBA)

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
		img:            i,
		tappedCallback: t,
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
// func (p *ImageColourPicker) Tapped(event *fyne.PointEvent) {
// 	// Get widths and heights
// 	imgWidth := float32(p.displayImage.Bounds().Dx())
// 	imgHeight := float32(p.displayImage.Bounds().Dy())
// 	canvasWidth := p.renderer.image.Size().Width
// 	canvasHeight := p.renderer.image.Size().Height

// 	// Calculate aspect ratios
// 	imgAspect := imgWidth / imgHeight
// 	canvasAspect := canvasWidth / canvasHeight

// 	// Get new width and height of image
// 	var newWidth, newHeight float32
// 	if imgAspect > canvasAspect {
// 		// Image is wider relative to canvas
// 		newWidth = canvasWidth
// 		newHeight = canvasWidth / imgAspect
// 	} else {
// 		// Image is taller relative to canvas
// 		newWidth = canvasHeight * imgAspect
// 		newHeight = canvasHeight
// 	}

// 	// Calculate bounds
// 	top := (canvasHeight / 2) - (newHeight / 2)
// 	bottom := (canvasHeight / 2) + (newHeight / 2)
// 	left := (canvasWidth / 2) - (newWidth / 2)
// 	right := (canvasWidth / 2) + (newWidth / 2)

// 	// fmt.Println("Clicked: ", event.Position)
// 	// fmt.Println("Canvas size: ", canvasWidth, canvasHeight)
// 	// fmt.Println("Canvas aspect: ", canvasAspect)
// 	// fmt.Println("Image size: ", imgWidth, imgHeight)
// 	// fmt.Println("Image aspect: ", imgAspect)
// 	// fmt.Println("New image size: ", newWidth, newHeight)
// 	// fmt.Printf("Top: %f, Bottom: %f, Left: %f, Right: %f\n", top, bottom, left, right)
// 	// fmt.Println("---")

// 	clickedX := event.Position.X
// 	clickedY := event.Position.Y

// 	// Check point within image
// 	if clickedX >= left && clickedY >= top && clickedX < right && clickedY < bottom {
// 		// Get clicked position relative to image location
// 		relativeClickedX := clickedX - left
// 		relativeClickedY := clickedY - top

// 		// Map clicked relative position to values in original image
// 		mappedClickedX := mapRange(relativeClickedX, 0, newWidth, 0, imgWidth)
// 		mappedClickedY := mapRange(relativeClickedY, 0, newHeight, 0, imgHeight)

// 		// fmt.Println(relativeClickedX, relativeClickedY, mappedClickedX, mappedClickedY)

// 		// Get NRGBA value a clicked point
// 		c := p.actualImage.At(mappedClickedX, mappedClickedY).(color.NRGBA)
// 		p.tappedCallback(c)
// 	}
// }

// Returns new renderer for ImageColourPicker
func (p *ImageColourPicker) CreateRenderer() fyne.WidgetRenderer {
	return &imageColourPickerRenderer{
		icp: p,
		img: p.img,
	}
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
