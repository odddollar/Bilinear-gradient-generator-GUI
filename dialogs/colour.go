package dialogs

import (
	"Bilinear-gradient-generator-GUI/global"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show selected colour information in dialog
func ShowColour(col color.NRGBA) {
	// Make image from colour
	img := canvas.NewImageFromImage(makeColour(col))
	img.FillMode = canvas.ImageFillOriginal

	// Set pixel channel values
	r := widget.NewEntry()
	r.SetText(fmt.Sprintf("%d", col.R))
	g := widget.NewEntry()
	g.SetText(fmt.Sprintf("%d", col.G))
	b := widget.NewEntry()
	b.SetText(fmt.Sprintf("%d", col.B))
	a := widget.NewEntry()
	a.SetText(fmt.Sprintf("%d", col.A))
	hex := widget.NewEntry()
	hex.SetText(fmt.Sprintf("#%02X%02X%02X", col.R, col.G, col.B))

	// Create label widgets and set styling
	rl := widget.NewLabel("Red")
	rl.Alignment = fyne.TextAlignTrailing
	rl.TextStyle = fyne.TextStyle{Bold: true}
	gl := widget.NewLabel("Green")
	gl.Alignment = fyne.TextAlignTrailing
	gl.TextStyle = fyne.TextStyle{Bold: true}
	bl := widget.NewLabel("Blue")
	bl.Alignment = fyne.TextAlignTrailing
	bl.TextStyle = fyne.TextStyle{Bold: true}
	al := widget.NewLabel("Alpha")
	al.Alignment = fyne.TextAlignTrailing
	al.TextStyle = fyne.TextStyle{Bold: true}
	hl := widget.NewLabel("Hex")
	hl.Alignment = fyne.TextAlignTrailing
	hl.TextStyle = fyne.TextStyle{Bold: true}

	// Create layout
	d := container.NewHBox(
		container.NewVBox(
			img,
			container.NewBorder(
				nil,
				nil,
				hl,
				nil,
				hex,
			),
		),
		container.NewGridWithColumns(
			2,
			rl,
			r,
			gl,
			g,
			bl,
			b,
			al,
			a,
		),
	)

	// Show dialog with layout
	dialog.ShowCustom(
		"Selected Colour",
		"OK",
		d,
		global.MainWindow,
	)
}

// Turn colour into full square with checkerboard
func makeColour(col color.NRGBA) image.Image {
	// Image parameters
	size := 128
	tileSize := 16

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: size, Y: size}
	colour := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	checkerboard := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	overlayed := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Iterate through pixels
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			// Set colour
			colour.SetNRGBA(x, y, col)

			// Determine the colour based on the current position
			if (x/tileSize+y/tileSize)%2 == 0 {
				// Set light grey pixel
				checkerboard.SetNRGBA(x, y, color.NRGBA{200, 200, 200, 255})
			} else {
				// Set white pixel
				checkerboard.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
		}
	}

	// Draw checkerboard base image
	draw.Draw(overlayed, overlayed.Bounds(), checkerboard, image.Point{}, draw.Src)

	// Draw colour over top
	draw.Draw(overlayed, overlayed.Bounds(), colour, image.Point{}, draw.Over)

	return overlayed
}
