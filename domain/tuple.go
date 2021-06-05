package domain

type Tuple struct {
	X  float64
	Y  float64
	Z  float64
	W  byte
}

func Vector(x, y, z float64) *Tuple{
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}

func Point(x, y, z float64) *Tuple{
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1,
	}
}

func (t Tuple) IsVector()  bool{
	return t.W == 0
}

func (t Tuple) IsPoint()  bool{
	return t.W == 1
}