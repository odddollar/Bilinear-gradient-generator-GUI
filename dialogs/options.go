package dialogs

import (
	"Bilinear-gradient-generator-GUI/global"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show options in dialog
func ShowOptions() {
	// Width and height entry boxes
	widthEntry := widget.NewEntry()
	widthEntry.Validator = validation.NewRegexp(`^([3-9]|[0-9]{2,})$`, "Must be greater than 2")
	widthEntry.SetText(strconv.Itoa(global.A.Preferences().IntWithFallback("width", 512)))
	heightEntry := widget.NewEntry()
	heightEntry.Validator = validation.NewRegexp(`^([3-9]|[0-9]{2,})$`, "Must be greater than 2")
	heightEntry.SetText(strconv.Itoa(global.A.Preferences().IntWithFallback("height", 512)))

	// Minimum alpha entry box
	// Anything less than 255 will give global.A random value for alpha between
	// the entered number and 255 inclusive
	alphaEntry := widget.NewEntry()
	alphaEntry.Validator = validation.NewRegexp(`^(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])$`, "Must be 0-255 inclusive")
	alphaEntry.SetText(strconv.Itoa(global.A.Preferences().IntWithFallback("minimumAlpha", 255)))

	// Hide corner button checkbox
	hideCorners := widget.NewCheck("", func(b bool) {})
	hideCorners.SetChecked(global.A.Preferences().Bool("hideCorners"))

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
				// Update width and height
				w, _ := strconv.Atoi(widthEntry.Text)
				global.A.Preferences().SetInt("width", w)
				h, _ := strconv.Atoi(heightEntry.Text)
				global.A.Preferences().SetInt("height", h)

				// Update minimum alpha
				t, _ := strconv.Atoi(alphaEntry.Text)
				global.A.Preferences().SetInt("minimumAlpha", t)

				// Update hide corners
				global.A.Preferences().SetBool("hideCorners", hideCorners.Checked)
				if hideCorners.Checked {
					global.MainWindow.SetContent(global.NoCornerButtonContent)
				} else {
					global.MainWindow.SetContent(global.CornerButtonContent)
				}
			}
		},
		global.MainWindow,
	)
	d.Resize(fyne.NewSize(360, 340))
	d.Show()
}
