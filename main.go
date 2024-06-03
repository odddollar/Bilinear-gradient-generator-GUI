package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create app and window
	a := app.New()
	w := a.NewWindow("Gradient Generator GUI")

	image := canvas.NewImageFromImage(generateImage())
	image.FillMode = canvas.ImageFillOriginal

	generateButton := widget.NewButton("Generate", func() {
		image.Image = generateImage()
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
