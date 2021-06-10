package models

func Color(red, green, blue float64) *Tuple {
	return &Tuple{red, green, blue, 1.0}
}

func (a Tuple) HadamardProduct(b Tuple) *Tuple {
	return &Tuple{a[X] * b[X], a[Y] * b[Y], a[Z] * b[Z], a[W] * b[W]}
}
