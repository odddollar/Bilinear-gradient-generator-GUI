package main

import (
	"image"
	"image/png"
	"os"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var imageCurrent image.Image

func main() {
	// Create app and window
	a := app.New()
	w := a.NewWindow("Bilinear Gradient Generator GUI")

	// Generate initial image
	generateImage()

	// Canvas to display current image
	imageDisplay := canvas.NewImageFromImage(imageCurrent)
	imageDisplay.FillMode = canvas.ImageFillOriginal

	// Button to generate new image
	generateButton := widget.NewButton("Generate", func() {
		generateImage()
		imageDisplay.Image = imageCurrent
		imageDisplay.Refresh()
	})
	generateButton.Importance = widget.HighImportance

	// Button to save current image
	saveButton := widget.NewButton("Save", func() {
		// Save image
		f, _ := os.Create("image.png")
		_ = png.Encode(f, imageCurrent)
	})

	// Create window layout
	w.SetContent(
		container.NewBorder(
			nil,
			container.NewBorder(
				nil,
				nil,
				nil,
				saveButton,
				generateButton,
			),
			nil,
			nil,
			imageDisplay,
		),
	)

	// Set window properties and run
	w.Show()
	a.Run()
}
