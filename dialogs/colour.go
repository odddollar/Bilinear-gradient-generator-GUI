package dialogs

import (
	"Bilinear-gradient-generator-GUI/global"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show selected colour information in dialog
func ShowColour(c color.NRGBA) {
	text := widget.NewLabel(fmt.Sprintf("Colour: %#v\n", c))

	dialog.ShowCustom(
		"Selected Colour",
		"OK",
		text,
		global.MainWindow,
	)
}
