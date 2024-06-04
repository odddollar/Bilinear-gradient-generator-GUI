package main

func refreshImage() {
	// Generate new image
	generateImage()

	// Update image display
	imageDisplay.Image = imageCurrent
	imageDisplay.Refresh()
}
