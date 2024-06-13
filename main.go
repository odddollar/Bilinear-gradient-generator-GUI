package main

import (
	"Bilinear-gradient-generator-GUI/global"
	"Bilinear-gradient-generator-GUI/widgets"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create app and window
	global.A = app.New()
	global.MainWindow = global.A.NewWindow("Bilinear Gradient Generator GUI")

	// Canvas to display current image
	global.ImageDisplay = canvas.NewImageFromImage(global.ImageCurrent)
	global.ImageDisplay.FillMode = canvas.ImageFillOriginal

	// Button to randomly generate new image
	global.RandomiseButton = widget.NewButton("Randomise", func() {
		randomiseCorners()
		refreshImage()
	})
	global.RandomiseButton.Importance = widget.HighImportance

	// Button to save current image
	global.SaveButton = widget.NewButton("Save PNG", saveImage)

	// Button to open options
	global.OptionsButton = widget.NewButtonWithIcon("", theme.MenuIcon(), showOptions)

	// Button to show about information
	global.AboutButton = widget.NewButtonWithIcon("", theme.InfoIcon(), showAbout)

	// Buttons to change corner pixel values
	global.TopLeftButton = widget.NewButton("...", func() { pickColour(&global.TopLeftPixel) })
	global.TopRightButton = widget.NewButton("...", func() { pickColour(&global.TopRightPixel) })
	global.BottomLeftButton = widget.NewButton("...", func() { pickColour(&global.BottomLeftPixel) })
	global.BottomRightButton = widget.NewButton("...", func() { pickColour(&global.BottomRightPixel) })

	// Create global.Spacer with same width as button with "..." text
	global.SpacerWidget = widgets.NewSpacer(widget.NewButton("...", func() {}).MinSize())

	// Create layout with corner buttons
	global.CornerButtonContent = container.NewBorder(
		container.NewBorder(
			nil,
			nil,
			global.TopLeftButton,
			global.TopRightButton,
			nil,
		),
		container.NewBorder(
			container.NewBorder(
				nil,
				nil,
				global.BottomLeftButton,
				global.BottomRightButton,
				nil,
			),
			nil,
			nil,
			container.NewHBox(
				global.SaveButton,
				global.OptionsButton,
				global.AboutButton,
			),
			global.RandomiseButton,
		),
		global.SpacerWidget,
		global.SpacerWidget,
		global.ImageDisplay,
	)

	// Create layout without corner buttons
	global.NoCornerButtonContent = container.NewBorder(
		nil,
		container.NewBorder(
			nil,
			nil,
			nil,
			container.NewHBox(
				global.SaveButton,
				global.OptionsButton,
				global.AboutButton,
			),
			global.RandomiseButton,
		),
		nil,
		nil,
		global.ImageDisplay,
	)

	// Generate initial image
	generateCheckerboard()
	randomiseCorners()
	refreshImage()

	// Set window layout based on options
	if global.A.Preferences().BoolWithFallback("hideCorners", false) {
		global.MainWindow.SetContent(global.NoCornerButtonContent)
	} else {
		global.MainWindow.SetContent(global.CornerButtonContent)
	}

	// Show window and run
	global.MainWindow.Show()
	global.A.Run()
}
