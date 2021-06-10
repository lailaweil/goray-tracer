package models

import (
	"fmt"
	"github.com/laiweil/goray-tracer/utils"
	"strings"
	"testing"
)

func TestCreateCanvas(t *testing.T) {
	cases := []struct {
		name   string
		w      int
		h      int
		result *Tuple
	}{
		{
			name: "Creating Canvas",
			w:    10,
			h:    20,
		},
	}

	for _, c := range cases {
		result := CreateCanvas(c.w, c.h)

		for _, p := range result.pixels {
			for _, pixel := range p {
				if !utils.FloatEquals(pixel[Red], 0) || !utils.FloatEquals(pixel[Green], 0) || !utils.FloatEquals(pixel[Blue], 0) {
					t.Errorf("%s - failed creating canvas: expected all black pixels but got %v", c.name, pixel)
				}
			}

		}

	}
}

func TestCanvas_WritePixel(t *testing.T) {
	cases := []struct {
		name   string
		canvas *Canvas
		x      int
		y      int
		color  *Tuple
		hasErr bool
	}{
		{
			name:   "Writing pixels to a canvas",
			canvas: CreateCanvas(10, 20),
			x:      2,
			y:      3,
			color:  Color(1, 0, 0),
			hasErr: false,
		},
		{
			name:   "Writing pixels to a canvas fails",
			canvas: CreateCanvas(10, 20),
			x:      10,
			y:      2,
			color:  Color(1, 0, 0),
			hasErr: true,
		},
		{
			name:   "Writing pixels to a canvas fails 2",
			canvas: CreateCanvas(10, 20),
			x:      3,
			y:      20,
			color:  Color(1, 0, 0),
			hasErr: true,
		},
	}

	for _, c := range cases {
		err := c.canvas.WritePixel(c.x, c.y, *c.color)

		if (err != nil && !c.hasErr) || (err == nil && c.hasErr) {
			t.Errorf("%s - failed Writing pixel: expected error to be (%v) but got (%v)", c.name, c.hasErr, err)
		}

		pixel, err := c.canvas.PixelAt(c.x, c.y)

		if err != nil && !c.hasErr {
			t.Errorf("%s - failed getting pixel: %v", c.name, err)
		}

		if err == nil && !pixel.Equal(*c.color) {
			t.Errorf("%s - failed Writing pixel: expected pixel to be (%v) but got (%v)", c.name, c.color, pixel)
		}

	}

}

func TestCanvas_GetPPMHeader(t *testing.T) {
	cases := []struct {
		name   string
		canvas *Canvas
		result string
	}{
		{
			name:   "Constructing the PPM header",
			canvas: CreateCanvas(5, 3),
			result: fmt.Sprintf(PPMHeader, PPMIdentifier, 5, 3, MaxColor),
		},
	}

	for _, c := range cases {
		result := c.canvas.GetPPMHeader()

		if strings.Compare(result, c.result) != 0 {
			t.Errorf("%s - failed creating PPM header of canvas: expected to be (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestCanvas_ToPPM(t *testing.T) {
	cases := []struct {
		name   string
		canvas *Canvas
		color1 *Tuple
		color2 *Tuple
		color3 *Tuple
	}{
		{
			name:   "Constructing the PPM pixel data",
			canvas: CreateCanvas(5, 3),
			color1: Color(1.5, 0, 0),
			color2: Color(0, 0.5, 0),
			color3: Color(-0.5, 0, 1),
		},
	}

	for _, c := range cases {
		c.canvas.WritePixel(0, 0, *c.color1)
		c.canvas.WritePixel(2, 1, *c.color2)
		c.canvas.WritePixel(4, 2, *c.color3)

		result := c.canvas.ToPPM()

		if !strings.Contains(result, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n") || !strings.Contains(result, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n") || !strings.Contains(result, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n") {
			t.Errorf("%s - failed creating PPM file of canvas: got (%v)", c.name, result)
		}

	}
}
