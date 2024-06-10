package main

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Holds global image state
var imageCurrent image.Image

// Holds corner pixels
var (
	topLeftPixel     image.NRGBA
	topRightPixel    image.NRGBA
	bottomLeftPixel  image.NRGBA
	bottomRightPixel image.NRGBA
)

// Variables to hold widgets
var (
	a                 fyne.App
	mainWindow        fyne.Window
	imageDisplay      *canvas.Image
	aboutButton       *widget.Button
	generateButton    *widget.Button
	saveButton        *widget.Button
	topLeftButton     *widget.Button
	topRightButton    *widget.Button
	bottomLeftButton  *widget.Button
	bottomRightButton *widget.Button
	spacer            *Spacer
)

func main() {
	// Generate initial image
	generateImage()

	// Create app and window
	a = app.New()
	mainWindow = a.NewWindow("Bilinear Gradient Generator GUI")

	// Canvas to display current image
	imageDisplay = canvas.NewImageFromImage(imageCurrent)
	imageDisplay.FillMode = canvas.ImageFillOriginal

	// Button to show about information
	aboutButton = widget.NewButtonWithIcon("", theme.InfoIcon(), showAbout)

	// Button to generate new image
	generateButton = widget.NewButton("Generate", refreshImage)
	generateButton.Importance = widget.HighImportance

	// Button to save current image
	saveButton = widget.NewButton("Save PNG", saveImage)

	// Buttons to change corner pixel values
	topLeftButton = widget.NewButton("...", func() {})
	topRightButton = widget.NewButton("...", func() {})
	bottomLeftButton = widget.NewButton("...", func() {})
	bottomRightButton = widget.NewButton("...", func() {})

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
				aboutButton,
			),
			generateButton,
		),
		spacer,
		spacer,
		imageDisplay,
	)

	// Set window properties and run
	mainWindow.SetContent(content)
	mainWindow.Show()
	a.Run()
}
