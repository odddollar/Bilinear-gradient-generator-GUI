package image_handling

import (
	"Bilinear-gradient-generator-GUI/global"
	"image"
	"image/color"
)

// Generate gradient from values in corner pixels and update image state
func generateImage() {
	// Convert corner pixel data type to individual channel arrays
	redArray := *createArray()
	greenArray := *createArray()
	blueArray := *createArray()
	alphaArray := *createArray()

	redArray[0][0] = global.TopLeftPixel.R
	greenArray[0][0] = global.TopLeftPixel.G
	blueArray[0][0] = global.TopLeftPixel.B
	alphaArray[0][0] = global.TopLeftPixel.A

	redArray[0][widthMinusOne] = global.TopRightPixel.R
	greenArray[0][widthMinusOne] = global.TopRightPixel.G
	blueArray[0][widthMinusOne] = global.TopRightPixel.B
	alphaArray[0][widthMinusOne] = global.TopRightPixel.A

	redArray[heightMinusOne][0] = global.BottomLeftPixel.R
	greenArray[heightMinusOne][0] = global.BottomLeftPixel.G
	blueArray[heightMinusOne][0] = global.BottomLeftPixel.B
	alphaArray[heightMinusOne][0] = global.BottomLeftPixel.A

	redArray[heightMinusOne][widthMinusOne] = global.BottomRightPixel.R
	greenArray[heightMinusOne][widthMinusOne] = global.BottomRightPixel.G
	blueArray[heightMinusOne][widthMinusOne] = global.BottomRightPixel.B
	alphaArray[heightMinusOne][widthMinusOne] = global.BottomRightPixel.A

	// Fill individual arrays with interpolated values
	fillArray(&redArray)
	fillArray(&greenArray)
	fillArray(&blueArray)
	fillArray(&alphaArray)

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}
	img := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Iterate through pixels
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Set colour
			red := redArray[y][x]
			green := greenArray[y][x]
			blue := blueArray[y][x]
			alpha := alphaArray[y][x]
			col := color.NRGBA{R: red, G: green, B: blue, A: alpha}
			img.SetNRGBA(x, y, col)
		}
	}

	global.ImageCurrent = img
}

// Calculate all interpolated values for array
func fillArray(array *[][]uint8) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			calculateAndSet(x, y, array)
		}
	}
}

// Calculate and update individual value for position in array
func calculateAndSet(posX, posY int, array *[][]uint8) {
	// Calculate weights with floating-point division
	topLeftWeight := float64((widthMinusOne-posX)*(heightMinusOne-posY)) / float64(widthMinusOne*heightMinusOne)
	topRightWeight := float64(posX*(heightMinusOne-posY)) / float64(widthMinusOne*heightMinusOne)
	bottomLeftWeight := float64((widthMinusOne-posX)*posY) / float64(widthMinusOne*heightMinusOne)
	bottomRightWeight := float64(posX*posY) / float64(widthMinusOne*heightMinusOne)

	// Calculate interpolated value
	calc := topLeftWeight*float64((*array)[0][0]) +
		topRightWeight*float64((*array)[0][widthMinusOne]) +
		bottomLeftWeight*float64((*array)[heightMinusOne][0]) +
		bottomRightWeight*float64((*array)[heightMinusOne][widthMinusOne])

	// Set value in array
	(*array)[posY][posX] = uint8(calc)
}
