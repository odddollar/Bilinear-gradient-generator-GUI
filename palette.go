package main

import (
	"image/color"

	"fyne.io/fyne/v2/dialog"
)

func pickColour(cornerColour *color.NRGBA) {
	// Open advanced colour picker dialog and set initial colour to current one
	d := dialog.NewColorPicker("Colour", "Pick corner colour", func(c color.Color) {
		// Convert color.Color to color.NRGBA
		(*cornerColour) = color.NRGBAModel.Convert(c).(color.NRGBA)

		// Refresh image with new corner value
		refreshImage()
	}, mainWindow)
	d.Advanced = true
	d.SetColor(cornerColour)
	d.Show()
}
