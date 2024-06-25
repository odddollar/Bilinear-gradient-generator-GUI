package widgets

import "math"

// Map value between old range to be between new range
func mapRange(value, oldMin, oldMax, newMin, newMax float32) int {
	mapped := (value-oldMin)*(newMax-newMin)/(oldMax-oldMin) + newMin
	return int(math.Round(float64(mapped)))
}
