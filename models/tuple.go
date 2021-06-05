package models

import (
	"github.com/laiweil/goray-tracer/utils"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W byte
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
		W: t.W * byte(scalar),
	}
}

func (t Tuple) Divide(scalar float64) *Tuple {
	return &Tuple{
		X: t.X / scalar,
		Y: t.Y / scalar,
		Z: t.Z / scalar,
		W: t.W / byte(scalar),
	}
}

func (t Tuple) Magnitude() float64 {
	return math.Sqrt(math.Exp2(t.X) + math.Exp2(t.Y) + math.Exp2(t.Z) + math.Exp2(float64(t.W)))
}
