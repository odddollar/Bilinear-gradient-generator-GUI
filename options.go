package main

import "fyne.io/fyne/v2/dialog"

// Show options in dialog
func showOptions() {
	// Create new dialog using form items
	dialog.ShowForm(
		"Options",
		"Save",
		"Cancel",
		nil,
		func(b bool) {

		},
		mainWindow,
	)
}
