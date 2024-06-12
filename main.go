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

	// Create window layout
	content := container.NewBorder(
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

	// Generate initial image
	generateCheckerboard()
	randomiseCorners()
	refreshImage()

	// Set window properties and run
	mainWindow.SetContent(content)
	mainWindow.Show()
	a.Run()
}
