package models

import (
	"fmt"
	"github.com/laiweil/goray-tracer/utils"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W int
}

func Vector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}

func Point(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1,
	}
}

func (t Tuple) IsVector() bool {
	return t.W == 0
}

func (t Tuple) IsPoint() bool {
	return t.W == 1
}

func (a Tuple) Equal(b Tuple) bool {
	result := utils.FloatEquals(a.X, b.X)
	fmt.Sprintf("hola resultado %t", result)
	return utils.FloatEquals(a.X, b.X) && utils.FloatEquals(a.Y, b.Y) && utils.FloatEquals(a.Z, b.Z) && a.W == b.W
}
func (a Tuple) Add(b Tuple) *Tuple {
	return &Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func (a Tuple) Sub(b Tuple) *Tuple {
	return &Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
	}
}

func (a *Tuple) Negate() *Tuple {
	return &Tuple{
		X: -a.X,
		Y: -a.Y,
		Z: -a.Z,
		W: -a.W,
	}
}

func (t Tuple) Multiply(scalar float64) *Tuple {
	return &Tuple{
		X: t.X * scalar,
		Y: t.Y * scalar,
		Z: t.Z * scalar,
		W: t.W * int(scalar),
	}
}

func (t Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W / int(scalar),
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2) + math.Pow(float64(t.W), 2))
}

func (t Tuple) Normalize() *Tuple {
	magnitude := t.Magnitude()
	return &Tuple{
		X: t.X / magnitude,
		Y: t.Y / magnitude,
		Z: t.Z / magnitude,
		W: t.W / int(magnitude),
	}
}

func (a Tuple) Dot(b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + float64(a.W*b.W)
}
