package image_handling

import "math/rand"

// Helper function to generate number within range
func randomRange(min, max int) uint8 {
	return uint8(rand.Intn(max-min) + min)
}
