package main

import (
	"Bilinear-gradient-generator-GUI/global"
	"image/color"

	"fyne.io/fyne/v2/dialog"
)

// Choose new colour for corner
func pickColour(cornerColour *color.NRGBA) {
	// Open advanced colour picker dialog and set initial colour to current one
	d := dialog.NewColorPicker("Colour", "Pick corner colour", func(c color.Color) {
		// Convert color.Color to color.NRGBA
		(*cornerColour) = color.NRGBAModel.Convert(c).(color.NRGBA)

		// Refresh image with new corner value
		refreshImage()
	}, global.MainWindow)
	d.Advanced = true
	d.SetColor(cornerColour)
	d.Show()
}
