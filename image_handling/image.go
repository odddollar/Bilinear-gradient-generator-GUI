package image_handling

import (
	"Bilinear-gradient-generator-GUI/global"
	"image"
	"image/color"
	"math/rand"
)

// Helper function to generate number within range
func randomRange(min, max int) uint8 {
	return uint8(rand.Intn(max-min) + min)
}

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
	// Generate new image from corner values
	generateImage()

	// Draw checkerboard then overlay image
	img := CombineCheckerboard()

	// Update image display
	global.ImageDisplay.Image = img
	global.ImageDisplay.Refresh()
}

// Generate gradient from values in corner pixels and update image state
func generateImage() {
	// Convert corner pixel data type to individual channel arrays
	redArray := [512][512]uint8{}
	greenArray := [512][512]uint8{}
	blueArray := [512][512]uint8{}
	alphaArray := [512][512]uint8{}

	redArray[0][0] = global.TopLeftPixel.R
	greenArray[0][0] = global.TopLeftPixel.G
	blueArray[0][0] = global.TopLeftPixel.B
	alphaArray[0][0] = global.TopLeftPixel.A

	redArray[0][511] = global.TopRightPixel.R
	greenArray[0][511] = global.TopRightPixel.G
	blueArray[0][511] = global.TopRightPixel.B
	alphaArray[0][511] = global.TopRightPixel.A

	redArray[511][0] = global.BottomLeftPixel.R
	greenArray[511][0] = global.BottomLeftPixel.G
	blueArray[511][0] = global.BottomLeftPixel.B
	alphaArray[511][0] = global.BottomLeftPixel.A

	redArray[511][511] = global.BottomRightPixel.R
	greenArray[511][511] = global.BottomRightPixel.G
	blueArray[511][511] = global.BottomRightPixel.B
	alphaArray[511][511] = global.BottomRightPixel.A

	// Fill individual arrays with interpolated values
	fillArray(&redArray)
	fillArray(&greenArray)
	fillArray(&blueArray)
	fillArray(&alphaArray)

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// Set colour
			red := uint8(redArray[y][x])
			green := uint8(greenArray[y][x])
			blue := uint8(blueArray[y][x])
			alpha := uint8(alphaArray[y][x])
			col := color.NRGBA{R: red, G: green, B: blue, A: alpha}
			img.SetNRGBA(x, y, col)
		}
	}

	global.ImageCurrent = img
}

// Calculate all interpolated values for array
func fillArray(array *[512][512]uint8) {
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			calculateAndSet(x, y, array)
		}
	}
}

// Calculate and update individual value for position in array
func calculateAndSet(posX, posY int, array *[512][512]uint8) {
	// Calculate weights with floating-point division
	topLeftWeight := float64((511-posX)*(511-posY)) / (511 * 511)
	topRightWeight := float64(posX*(511-posY)) / (511 * 511)
	bottomLeftWeight := float64((511-posX)*posY) / (511 * 511)
	bottomRightWeight := float64(posX*posY) / (511 * 511)

	// Calculate interpolated value
	calc := topLeftWeight*float64(array[0][0]) +
		topRightWeight*float64(array[0][511]) +
		bottomLeftWeight*float64(array[511][0]) +
		bottomRightWeight*float64(array[511][511])

	// Set value in array
	(*array)[posY][posX] = uint8(calc)
}
