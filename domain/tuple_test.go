package domain

import "testing"

func TestPoint(t *testing.T) {

	cases := []struct{
		name string
		x float64
		y float64
		z float64
	}{
		{
			name:    "OK/Point",
			x:       4,
			y:       -4,
			z:       3,
		},
	}

	for _, c := range cases {
		result := Point(c.x, c.y, c.z)

		if result.W != 1{
			t.Errorf("%s - failed creating point: expected w to be (%v) but got (%v)", c.name, 1, result)
		}
	}
}


func TestVector(t *testing.T) {

	cases := []struct{
		name string
		x float64
		y float64
		z float64
	}{
		{
			name:    "OK/Vector",
			x:       4,
			y:       -4,
			z:       3,
		},
	}

	for _, c := range cases {
		result := Vector(c.x, c.y, c.z)

		if result.W != 0{
			t.Errorf("%s - failed creating point: expected w to be (%v) but got (%v)", c.name, 0, result)
		}
	}
}


func TestTuple_IsPoint(t *testing.T) {

	cases := []struct{
		name string
		point *Tuple
		isPoint bool
	}{
		{
			name:    "OK/Point",
			point:   &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 1.0,
			},
			isPoint: true,
		},
		{
			name:    "OK/not_Point",
			point:   &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 0,
			},
			isPoint: false,
		},
	}

	for _, c := range cases {
		result := c.point.IsPoint()

		if result != c.isPoint{
			t.Errorf("%s - failed checking for point: expected (%v) but got (%v)", c.name, c.isPoint, result)
		}
	}
}


func TestTuple_IsVector(t *testing.T) {

	cases := []struct{
		name string
		point *Tuple
		isVector bool
	}{
		{
			name:    "OK/Vector",
			point:   &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 0,
			},
			isVector: true,
		},
		{
			name:    "OK/not_Vector",
			point:   &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 1.0,
			},
			isVector: false,
		},
	}

	for _, c := range cases {
		result := c.point.IsVector()

		if result != c.isVector{
			t.Errorf("%s - failed checking for point: expected (%v) but got (%v)", c.name, c.isVector, result)
		}
	}
}


