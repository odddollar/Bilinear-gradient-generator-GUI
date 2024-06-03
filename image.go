package main

import (
	"image"
	"image/color"
	"math/rand"
)

func generateImage() image.Image {
	// Create colour arrays with random corner values
	redArray := [512][512]int{}
	greenArray := [512][512]int{}
	blueArray := [512][512]int{}

	redArray[0][0] = rand.Intn(255)
	greenArray[0][0] = rand.Intn(255)
	blueArray[0][0] = rand.Intn(255)

	redArray[0][511] = rand.Intn(255)
	greenArray[0][511] = rand.Intn(255)
	blueArray[0][511] = rand.Intn(255)

	redArray[511][0] = rand.Intn(255)
	greenArray[511][0] = rand.Intn(255)
	blueArray[511][0] = rand.Intn(255)

	redArray[511][511] = rand.Intn(255)
	greenArray[511][511] = rand.Intn(255)
	blueArray[511][511] = rand.Intn(255)

	// Fill individual arrays with interpolated values
	fillArray(&redArray)
	fillArray(&greenArray)
	fillArray(&blueArray)

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// Set colour
			red := uint8(redArray[y][x])
			green := uint8(greenArray[y][x])
			blue := uint8(blueArray[y][x])
			col := color.RGBA{R: red, G: green, B: blue, A: 0xff}
			img.Set(x, y, col)
		}
	}

	return img
}

func fillArray(array *[512][512]int) {
	// Top Row
	for x := 1; x < 511; x++ {
		(*array)[0][x] = getValue(x, 0, array)
	}
	// Middle chunk
	for y := 1; y < 511; y++ {
		for x := 0; x < 512; x++ {
			(*array)[y][x] = getValue(x, y, array)
		}
	}
	// Bottom row
	for x := 1; x < 511; x++ {
		(*array)[511][x] = getValue(x, 511, array)
	}
}

func getValue(posX, posY int, array *[512][512]int) int {
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

	return int(calc)
}
