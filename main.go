package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Holds global image state
var (
	imageCurrent image.Image
	checkerboard image.Image
)

// Holds corner pixels
var (
	topLeftPixel     color.NRGBA
	topRightPixel    color.NRGBA
	bottomLeftPixel  color.NRGBA
	bottomRightPixel color.NRGBA
)

// Variables to hold widgets
var (
	a                 fyne.App
	mainWindow        fyne.Window
	imageDisplay      *canvas.Image
	randomiseButton   *widget.Button
	saveButton        *widget.Button
	optionsButton     *widget.Button
	aboutButton       *widget.Button
	topLeftButton     *widget.Button
	topRightButton    *widget.Button
	bottomLeftButton  *widget.Button
	bottomRightButton *widget.Button
	spacer            *Spacer
)

// Variables to hold different layouts
// One holds layout with corner buttons, the other without
var (
	cornerButtonContent   *fyne.Container
	noCornerButtonContent *fyne.Container
)

func main() {
	// Create app and window
	a = app.New()
	mainWindow = a.NewWindow("Bilinear Gradient Generator GUI")

	// Canvas to display current image
	imageDisplay = canvas.NewImageFromImage(imageCurrent)
	imageDisplay.FillMode = canvas.ImageFillOriginal

	// Button to randomly generate new image
	randomiseButton = widget.NewButton("Randomise", func() {
		randomiseCorners()
		refreshImage()
	})
	randomiseButton.Importance = widget.HighImportance

	// Button to save current image
	saveButton = widget.NewButton("Save PNG", saveImage)

	// Button to open options
	optionsButton = widget.NewButtonWithIcon("", theme.MenuIcon(), showOptions)

	// Button to show about information
	aboutButton = widget.NewButtonWithIcon("", theme.InfoIcon(), showAbout)

	// Buttons to change corner pixel values
	topLeftButton = widget.NewButton("...", func() { pickColour(&topLeftPixel) })
	topRightButton = widget.NewButton("...", func() { pickColour(&topRightPixel) })
	bottomLeftButton = widget.NewButton("...", func() { pickColour(&bottomLeftPixel) })
	bottomRightButton = widget.NewButton("...", func() { pickColour(&bottomRightPixel) })

	// Create spacer with same width as button with "..." text
	spacer = NewSpacer(widget.NewButton("...", func() {}).MinSize())

	// Create layout with corner buttons
	cornerButtonContent = container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			topLeftButton,
			topRightButton,
			nil,
		),
		container.NewBorder(
			container.NewBorder(
				nil,
				nil,
				bottomLeftButton,
				bottomRightButton,
				nil,
			),
			nil,
			nil,
			container.NewHBox(
				saveButton,
				optionsButton,
				aboutButton,
			),
			randomiseButton,
		),
		spacer,
		spacer,
		imageDisplay,
	)

	// Create layout without corner buttons
	noCornerButtonContent = container.NewBorder(
		nil,
		container.NewBorder(
			nil,
			nil,
			nil,
			container.NewHBox(
				saveButton,
				optionsButton,
				aboutButton,
			),
			randomiseButton,
		),
		nil,
		nil,
		imageDisplay,
	)

	// Generate initial image
	generateCheckerboard()
	randomiseCorners()
	refreshImage()

	// Set window layout based on options
	if a.Preferences().BoolWithFallback("hideCorners", false) {
		mainWindow.SetContent(noCornerButtonContent)
	} else {
		mainWindow.SetContent(cornerButtonContent)
	}

	// Show window and run
	mainWindow.Show()
	a.Run()
}
