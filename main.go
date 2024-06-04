package main

import (
	"image"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var imgState image.Image

func main() {
	// Create app and window
	a := app.New()
	w := a.NewWindow("Bilinear Gradient Generator GUI")

	// Generate initial image
	generateImage()

	image := canvas.NewImageFromImage(imgState)
	image.FillMode = canvas.ImageFillOriginal

	generateButton := widget.NewButton("Generate", func() {
		generateImage()
		image.Image = imgState
		image.Refresh()
	})
	generateButton.Importance = widget.HighImportance

	// Create window layout
	w.SetContent(
		container.NewVBox(
			image,
			generateButton,
		),
	)

	// Set window properties and run
	w.Show()
	a.Run()
}
