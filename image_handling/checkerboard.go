package image_handling

import (
	"Bilinear-gradient-generator-GUI/global"
	"image"
	"image/color"
	"image/draw"
)

// Generate checkerboard and overlay main image
func combineCheckerboard() image.Image {
	// Size of individual tiles
	tileSize := 16

	// Create image bounding box
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: width, Y: height}
	checkerboard := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})
	overlayed := image.NewNRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Determine the colour based on the current position
			if (x/tileSize+y/tileSize)%2 == 0 {
				// Set light grey pixel
				checkerboard.SetNRGBA(x, y, color.NRGBA{200, 200, 200, 255})
			} else {
				// Set white pixel
				checkerboard.SetNRGBA(x, y, color.NRGBA{255, 255, 255, 255})
			}
		}
	}

	// Draw checkerboard base image
	draw.Draw(overlayed, overlayed.Bounds(), checkerboard, image.Point{}, draw.Src)

	// Draw gradient over top
	draw.Draw(overlayed, overlayed.Bounds(), global.ImageCurrent, image.Point{}, draw.Over)

	return overlayed
}
