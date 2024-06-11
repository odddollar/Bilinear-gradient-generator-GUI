package main

import (
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

// Save current image state (without checkerboard)
func saveImage() {
	// Create new dialog to save image
	d := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
		// Prevent crashing if nothing was selected
		if uc == nil {
			return
		}
		defer uc.Close()

		// Encode and save image
		png.Encode(uc, imageCurrent)

		// Successful save
		dialog.ShowInformation("Success", "Save successful", mainWindow)
	}, mainWindow)
	d.SetFileName("image.png")
	d.Show()
}
