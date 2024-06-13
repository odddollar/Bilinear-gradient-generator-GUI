package image_handling

import (
	"Bilinear-gradient-generator-GUI/global"
	"image/color"
	"math/rand"
)

var width, height, widthMinusOne, heightMinusOne int

// Randomise values in all corner pixel variables
func RandomiseCorners() {
	// What minimum alpha should be used
	minAlpha := global.A.Preferences().IntWithFallback("minimumAlpha", 255)

	// Randomise each corner
	global.TopLeftPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: randomRange(minAlpha, 256),
	}
	global.TopRightPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: randomRange(minAlpha, 256),
	}
	global.BottomLeftPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: randomRange(minAlpha, 256),
	}
	global.BottomRightPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: randomRange(minAlpha, 256),
	}
}

// Refresh displayed image
func RefreshImage() {
	// Get width and height from preferences
	width = global.A.Preferences().IntWithFallback("width", 512)
	height = global.A.Preferences().IntWithFallback("height", 512)
	widthMinusOne = width - 1
	heightMinusOne = height - 1

	// Generate new image from corner values
	generateImage()

	// Draw checkerboard then overlay image
	img := CombineCheckerboard()

	// Update image display
	global.ImageDisplay.Image = img
	global.ImageDisplay.Refresh()
}
