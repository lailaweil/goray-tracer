package models

import (
	"errors"
	genericError "github.com/laiweil/goray-tracer/errors"
)

type Canvas struct {
	width  int
	height int
	pixels [][]Tuple
}

func CreateCanvas(width int, height int) *Canvas {
	pixels := make([][]Tuple, width)
	rows := make([]Tuple, width*height)
	for i := 0; i < width; i++ {
		pixels[i] = rows[i*height : (i+1)*height]
	}

	return &Canvas{
		width:  width,
		height: height,
		pixels: pixels,
	}
}

func (c Canvas) Height() int {
	return c.height
}

func (c Canvas) Width() int {
	return c.width
}

func (c *Canvas) WritePixel(x, y int, color Tuple) error {
	if x >= c.width || y >= c.height {
		return genericError.CreateGenericError(errors.New("pixels out of canvas"), "Error: pixel out of canvas")
	}
	c.pixels[x][y] = color
	return nil
}

func (c *Canvas) PixelAt(x, y int) (*Tuple, error) {
	if x >= c.width || y >= c.height {
		return nil, genericError.CreateGenericError(errors.New("pixels out of canvas"), "Error: pixel out of canvas")
	}
	return &c.pixels[x][y], nil
}
