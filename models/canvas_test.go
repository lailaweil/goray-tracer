package models

import (
	"github.com/laiweil/goray-tracer/utils"
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
