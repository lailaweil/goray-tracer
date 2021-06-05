package utils

const (
	epsilon = 0.0001
)

func FloatEquals(a float64, b float64)  bool{
	return abs(a - b) < epsilon
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}