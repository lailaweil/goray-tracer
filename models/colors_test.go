package models

import (
	"fmt"
	"github.com/laiweil/goray-tracer/utils"
	"testing"
)

func TestColor(t *testing.T) {

	cases := []struct {
		name  string
		red   float64
		green float64
		blue  float64
	}{
		{
			name:  "OK/Color",
			red:   -0.5,
			green: 0.4,
			blue:  1.7,
		},
	}

	for _, c := range cases {
		result := Color(c.red, c.green, c.blue)

		if !utils.FloatEquals(result[W], 1.0) || !utils.FloatEquals(result[Red], -0.5) || !utils.FloatEquals(result[Blue], 1.7) {
			t.Errorf("%s - failed creating color: expected w to be (%v) but got (%v)", c.name, fmt.Sprintf("(%f, %f, %f", c.red, c.green, c.blue), result)
		}
	}
}

func TestTuple_HadamardProduct(t *testing.T) {

	cases := []struct {
		name   string
		color1 *Tuple
		color2 *Tuple
		result *Tuple
	}{
		{
			name:   "Blending Colors",
			color1: Color(1, 0.2, 0.4),
			color2: Color(0.9, 1, 0.1),
			result: Color(0.9, 0.2, 0.04),
		},
	}

	for _, c := range cases {
		result := c.color1.HadamardProduct(*c.color2)

		if !result.Equal(*c.result) {
			t.Errorf("%s - failed blending colors: expected w to be (%v) but got (%v)", c.name, c.result, result)
		}
	}
}
