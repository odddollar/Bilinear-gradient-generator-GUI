package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show options in dialog
func showOptions() {
	// Width and height entry boxes
	widthEntry := widget.NewEntry()
	widthEntry.Validator = validation.NewRegexp(`^([3-9]|[0-9]{2,})$`, "Must be greater than 2")
	heightEntry := widget.NewEntry()
	heightEntry.Validator = validation.NewRegexp(`^([3-9]|[0-9]{2,})$`, "Must be greater than 2")

	widthEntry.SetText("512")
	heightEntry.SetText("512")

	// Minimum alpha entry box
	// Anything less than 255 will give a random value for alpha between
	// the entered number and 255 inclusive
	alphaEntry := widget.NewEntry()
	alphaEntry.Validator = validation.NewRegexp(`^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])$`, "Must be 0-255 inclusive")
	alphaEntry.SetText(strconv.Itoa(a.Preferences().IntWithFallback("minimumAlpha", 255)))

	// Hide corner button checkbox
	hideCorners := widget.NewCheck("", func(b bool) {})

	// Create options layout
	options := []*widget.FormItem{
		{Text: "Image width", Widget: widthEntry, HintText: "In pixels"},
		{Text: "Image height", Widget: heightEntry, HintText: "In pixels"},
		{Text: "Minimum alpha", Widget: alphaEntry, HintText: "Minimum random alpha value"},
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
				// Update minimum alpha
				t, _ := strconv.Atoi(alphaEntry.Text)
				a.Preferences().SetInt("minimumAlpha", t)
			}
		},
		mainWindow,
	)
	d.Resize(fyne.NewSize(360, 340))
	d.Show()
}
