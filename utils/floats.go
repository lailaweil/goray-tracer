package utils

import "math"

const (
	epsilon = 0.0001
)

func FloatEquals(a float64, b float64) bool {
	return math.Abs(a-b) < epsilon
}
