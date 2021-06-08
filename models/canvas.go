package models

import (
	"errors"
	"fmt"
	genericError "github.com/laiweil/goray-tracer/errors"
	"io"
	"math"
)

type Canvas struct {
	width  int
	height int
	pixels [][]Tuple
}

func CreateCanvas(width int, height int) *Canvas {
	pixels := make([][]Tuple, height)
	rows := make([]Tuple, width*height)
	for i := 0; i < height; i++ {
		pixels[i] = rows[i*width : (i+1)*width]
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

func (c Canvas) Pixels() [][]Tuple {
	return c.pixels
}

func (c *Canvas) WritePixel(x, y int, color Tuple) error {
	if x >= c.width || y >= c.height {
		return genericError.CreateGenericError(errors.New("pixels out of canvas"), "Error: pixel out of canvas")
	}
	c.pixels[y][x] = color
	return nil
}

func (c *Canvas) PixelAt(x, y int) (*Tuple, error) {
	if x >= c.width || y >= c.height {
		return nil, genericError.CreateGenericError(errors.New("pixels out of canvas"), "Error: pixel out of canvas")
	}
	return &c.pixels[y][x], nil
}

func (c *Canvas) WritePPM(w io.Writer) error {
	ppm := c.ToPPM()

	_, err := fmt.Fprintln(w, ppm)

	if err != nil {
		return genericError.CreateGenericError(err, "Error writing to PPM file")
	}

	return nil
}

func (c Canvas) ToPPM() string {
	finalPPM := c.GetPPMHeader()
	rowPixels := ""

	for _, pixels := range c.pixels {
		for _, pixel := range pixels {
			rowPixels = fmt.Sprintf("%s %d %d %d", rowPixels, LimitRGB(pixel[Red]), LimitRGB(pixel[Green]), LimitRGB(pixel[Blue]))
		}
		finalPPM = fmt.Sprintf("%s\n%s", finalPPM, rowPixels)
		rowPixels = ""
	}

	return fmt.Sprintf("%s\n", finalPPM)
}

func LimitRGB(colorComponent float64) int {

	if colorComponent > 1 {
		return MaxColor
	} else if colorComponent < 0 {
		return 0
	}

	return int(math.Round(255 * colorComponent))
}

func (c *Canvas) GetPPMHeader() string {
	return fmt.Sprintf(PPMHeader, PPMIdentifier, c.width, c.height, MaxColor)
}
