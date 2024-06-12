package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show options in dialog
func showOptions() {
	// Width and height entry boxes
	widthEntry := widget.NewEntry()
	widthEntry.Validator = validation.NewRegexp(`^[3-9]|[0-9]{2,}$`, "Must be greater than 2")
	heightEntry := widget.NewEntry()
	heightEntry.Validator = validation.NewRegexp(`^[3-9]|[0-9]{2,}$`, "Must be greater than 2")

	// Randomise alpha checkbox
	alphaCheckbox := widget.NewCheck("", func(b bool) {})
	alphaCheckbox.SetChecked(a.Preferences().BoolWithFallback("randomiseAlpha", false))

	// Hide corner button checkbox
	hideCorners := widget.NewCheck("", func(b bool) {})

	// Create options layout
	options := []*widget.FormItem{
		{Text: "Image width", Widget: widthEntry, HintText: "In pixels"},
		{Text: "Image height", Widget: heightEntry, HintText: "In pixels"},
		{Text: "Randomise alpha", Widget: alphaCheckbox},
		{Text: "Hide corner buttons", Widget: hideCorners},
	}

	// Create new dialog using form items
	d := dialog.NewForm(
		"Options",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				// Update alpha
				a.Preferences().SetBool("randomiseAlpha", alphaCheckbox.Checked)
			}
		},
		mainWindow,
	)
	d.Resize(fyne.NewSize(350, 315))
	d.Show()
}
