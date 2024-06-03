package main

import (
	"image"
	"image/color"
	"math/rand"
)

func generateImage() image.Image {
	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Iterate through pixels
	for y := 0; y < 512; y++ {
		for x := 0; x < 512; x++ {
			// Set colour
			red := uint8(rand.Intn(255))
			green := uint8(rand.Intn(255))
			blue := uint8(rand.Intn(255))
			col := color.RGBA{R: red, G: green, B: blue, A: 0xff}
			img.Set(x, y, col)
		}
	}

	return img
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
