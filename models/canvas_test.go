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
