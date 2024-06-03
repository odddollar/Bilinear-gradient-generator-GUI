package main

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// set colour
			red := uint8(255)
			green := uint8(200)
			blue := uint8(255)
			col := color.RGBA{R: red, G: green, B: blue, A: 0xff}
			img.Set(x, y, col)
		}
	}

	image := canvas.NewImageFromImage(img)
	image.FillMode = canvas.ImageFillOriginal

	w.SetContent(image)
	w.Show()
	a.Run()
}
