package image_handling

import (
	"Bilinear-gradient-generator-GUI/global"
	"image"
	"image/color"
	"image/draw"
)

// Generate checkerboard and update checkerboard state
func GenerateCheckerboard() {
	// Size of image and individual tiles
	width := 512
	height := 512
	tileSize := 16

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Determine the colour based on the current position
			if (x/tileSize+y/tileSize)%2 == 0 {
				// Set light grey pixel
				img.SetNRGBA(x, y, color.NRGBA{200, 200, 200, 255})
			} else {
				// Set white pixel
				img.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
		}
	}

	global.Checkerboard = img
}

// Overlay image state on checkerboard and return new image
func CombineCheckerboard() image.Image {
	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 512, Y: 512}
	img := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	// Draw checkerboard base image
	draw.Draw(img, img.Bounds(), global.Checkerboard, image.Point{}, draw.Src)

	// Draw gradient over top
	draw.Draw(img, img.Bounds(), global.ImageCurrent, image.Point{}, draw.Over)

	return img
}
