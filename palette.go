package main

import (
	"image/color"

	"fyne.io/fyne/v2/dialog"
)

func pickColour(cornerColour *color.NRGBA) {
	// Open advanced colour picker dialog and set initial colour to current one
	d := dialog.NewColorPicker("Colour", "Pick corner colour", func(c color.Color) {

	}, mainWindow)
	d.Advanced = true
	d.SetColor(cornerColour)
	d.Show()
}
