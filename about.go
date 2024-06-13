package main

import (
	"Bilinear-gradient-generator-GUI/global"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show about information in dialog
func showAbout() {
	// Create layout
	// Separate markdown widget for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("Version: **%s**", global.A.Metadata().Version)),
		widget.NewRichTextFromMarkdown("Created by: [odddollar (Simon Eason)](https://github.com/odddollar)"),
		widget.NewRichTextFromMarkdown("Source: [github.com/odddollar/Bilinear-gradient-generator-GUI](https://github.com/odddollar/Bilinear-gradient-generator-GUI)"),
	)

	// Show information dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		global.MainWindow,
	)
}
