package main

import (
	"image"
	"image/color"
	"math/rand"
)

func randomiseCorners() {
	// Randomise each corner
	topLeftPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
	topRightPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
	bottomLeftPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
	bottomRightPixel = color.NRGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func refreshImage() {
	// Generate new image from corner values
	generateImage()

	// Update image display
	imageDisplay.Image = imageCurrent
	imageDisplay.Refresh()
}

func generateImage() {
	// Convert corner pixel data type to individual channel arrays
	redArray := [512][512]uint8{}
	greenArray := [512][512]uint8{}
	blueArray := [512][512]uint8{}

	redArray[0][0] = topLeftPixel.R
	greenArray[0][0] = topLeftPixel.G
	blueArray[0][0] = topLeftPixel.B

	redArray[0][511] = topRightPixel.R
	greenArray[0][511] = topRightPixel.G
	blueArray[0][511] = topRightPixel.B

	redArray[511][0] = bottomLeftPixel.R
	greenArray[511][0] = bottomLeftPixel.G
	blueArray[511][0] = bottomLeftPixel.B

	redArray[511][511] = bottomRightPixel.R
	greenArray[511][511] = bottomRightPixel.G
	blueArray[511][511] = bottomRightPixel.B

	// Fill individual arrays with interpolated values
	fillArray(&redArray)
	fillArray(&greenArray)
	fillArray(&blueArray)

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
			col := color.NRGBA{R: red, G: green, B: blue, A: 0xff}
			img.SetNRGBA(x, y, col)
		}
	}

	imageCurrent = img
}

func fillArray(array *[512][512]uint8) {
	// Calculate all interpolated values for the array
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			calculateAndSet(x, y, array)
		}
	}
}

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
