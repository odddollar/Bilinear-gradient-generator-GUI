package main

import (
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func saveImage() {
	// Create new dialog to save image
	d := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		// Prevent crashing if nothing was selected
		if uc == nil {
			return
		}
		defer uc.Close()

		// TODO: Ensure file has .png extension

		// Encode and save image
		png.Encode(uc, imageCurrent)
	}, mainWindow)
	d.Show()
}
