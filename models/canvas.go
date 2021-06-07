package models

type Canvas struct {
	Width  int
	Height int
	pixels [][]Tuple
}

func CreateCanvas(width int, height int) *Canvas {
	pixels := make([][]Tuple, width)
	rows := make([]Tuple, width*height)
	for i := 0; i < width; i++ {
		pixels[i] = rows[i*height : (i+1)*height]
	}

	return &Canvas{
		Width:  width,
		Height: height,
		pixels: pixels,
	}
}
