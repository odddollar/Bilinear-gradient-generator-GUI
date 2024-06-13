package image_handling

import "math/rand"

// Helper function to generate number within range
func randomRange(min, max int) uint8 {
	return uint8(rand.Intn(max-min) + min)
}

// Initialises array with image's width and height
func createArray() *[][]uint8 {
	matrix := make([][]uint8, height)
	for i := range matrix {
		matrix[i] = make([]uint8, width)
	}
	return &matrix
}
