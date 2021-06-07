package models

import (
	"github.com/laiweil/goray-tracer/utils"
	"math"
)

type Tuple [4]float64

func Vector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0.0}
}

func Point(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1.0}
}

func Color(red, green, blue float64) *Tuple {
	return &Tuple{red, green, blue, 1.0}
}

func (t Tuple) IsVector() bool {
	return utils.FloatEquals(t[W], 0)
}

func (t Tuple) IsPoint() bool {
	return utils.FloatEquals(t[W], 1)
}

func (a Tuple) Equal(b Tuple) bool {
	return utils.FloatEquals(a[X], b[X]) && utils.FloatEquals(a[Y], b[Y]) && utils.FloatEquals(a[Z], b[Z]) && utils.FloatEquals(a[W], b[W])
}
func (a Tuple) Add(b Tuple) *Tuple {
	return &Tuple{
		a[X] + b[X],
		a[Y] + b[Y],
		a[Z] + b[Z],
		a[W] + b[W],
	}
}

func (a Tuple) Sub(b Tuple) *Tuple {
	return &Tuple{
		X: a[X] - b[X],
		Y: a[Y] - b[Y],
		Z: a[Z] - b[Z],
		W: a[W] - b[W],
	}
}

func (a *Tuple) Negate() *Tuple {
	return &Tuple{
		-a[X],
		-a[Y],
		-a[Z],
		-a[W],
	}
}

func (t Tuple) Multiply(scalar float64) *Tuple {
	return &Tuple{
		t[X] * scalar,
		t[Y] * scalar,
		t[Z] * scalar,
		t[W] * scalar,
	}
}

func (t Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		X: t[X] / scalar,
		Y: t[Y] / scalar,
		Z: t[Z] / scalar,
		W: t[W] / scalar,
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t[X], 2) + math.Pow(t[Y], 2) + math.Pow(t[Z], 2) + math.Pow(t[W], 2))
}

func (t Tuple) Normalize() *Tuple {
	magnitude := t.Magnitude()
	return &Tuple{
		t[X] / magnitude,
		t[Y] / magnitude,
		t[Z] / magnitude,
		t[W] / magnitude,
	}
}

func (a Tuple) Dot(b Tuple) float64 {
	return a[X]*b[X] + a[Y]*b[Y] + a[Z]*b[Z] + a[W]*b[W]
}

func (a Tuple) Cross(b Tuple) *Tuple {
	return Vector(a[Y]*b[Z]-a[Z]*b[Y], a[Z]*b[X]-a[X]*b[Z], a[X]*b[Y]-a[Y]*b[X])
}
