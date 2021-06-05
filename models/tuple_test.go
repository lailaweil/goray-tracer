package models

import (
	"github.com/laiweil/goray-tracer/utils"
	"math"
	"testing"
)

func TestPoint(t *testing.T) {

	cases := []struct {
		name string
		x    float64
		y    float64
		z    float64
	}{
		{
			name: "OK/Point",
			x:    4,
			y:    -4,
			z:    3,
		},
	}

	for _, c := range cases {
		result := Point(c.x, c.y, c.z)

		if result.W != 1 {
			t.Errorf("%s - failed creating point: expected w to be (%v) but got (%v)", c.name, 1, result)
		}
	}
}

func TestVector(t *testing.T) {

	cases := []struct {
		name string
		x    float64
		y    float64
		z    float64
	}{
		{
			name: "OK/Vector",
			x:    4,
			y:    -4,
			z:    3,
		},
	}

	for _, c := range cases {
		result := Vector(c.x, c.y, c.z)

		if result.W != 0 {
			t.Errorf("%s - failed creating point: expected w to be (%v) but got (%v)", c.name, 0, result)
		}
	}
}

func TestTuple_IsPoint(t *testing.T) {

	cases := []struct {
		name    string
		point   *Tuple
		isPoint bool
	}{
		{
			name: "OK/Point",
			point: &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 1.0,
			},
			isPoint: true,
		},
		{
			name: "OK/not_Point",
			point: &Tuple{
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

		if result != c.isPoint {
			t.Errorf("%s - failed checking for point: expected (%v) but got (%v)", c.name, c.isPoint, result)
		}
	}
}

func TestTuple_IsVector(t *testing.T) {

	cases := []struct {
		name     string
		point    *Tuple
		isVector bool
	}{
		{
			name: "OK/Vector",
			point: &Tuple{
				X: 4.3,
				Y: -4.2,
				Z: 3.1,
				W: 0,
			},
			isVector: true,
		},
		{
			name: "OK/not_Vector",
			point: &Tuple{
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

		if result != c.isVector {
			t.Errorf("%s - failed checking for point: expected (%v) but got (%v)", c.name, c.isVector, result)
		}
	}
}

func TestTuple_Equal(t *testing.T) {

}

func TestTuple_Add(t *testing.T) {
	cases := []struct {
		name   string
		tupleA *Tuple
		tupleB *Tuple
		result *Tuple
	}{
		{
			name:   "OK/Addition_Point_and_Vector",
			tupleA: Point(3, -2, 5),
			tupleB: Vector(-2, 3, 1),
			result: Point(1, 1, 6),
		},
		{
			name:   "OK/Addition_Vectors",
			tupleA: Vector(3, -2, 5),
			tupleB: Vector(-2, 3, 1),
			result: Vector(1, 1, 6),
		},
	}

	for _, c := range cases {
		result := c.tupleA.Add(*c.tupleB)

		if !c.result.Equal(*result) {
			t.Errorf("%s - failed adding tuple a to tuple b: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Sub(t *testing.T) {
	cases := []struct {
		name   string
		tupleA *Tuple
		tupleB *Tuple
		result *Tuple
	}{
		{
			name:   "OK/Subtraction_Points",
			tupleA: Point(3, 2, 1),
			tupleB: Point(5, 6, 7),
			result: Vector(-2, -4, -6),
		},
		{
			name:   "OK/Subtraction_Vector_from_Point",
			tupleA: Point(3, 2, 1),
			tupleB: Vector(5, 6, 7),
			result: Point(-2, -4, -6),
		},
		{
			name:   "OK/Subtraction_Vectors",
			tupleA: Vector(3, 2, 1),
			tupleB: Vector(5, 6, 7),
			result: Vector(-2, -4, -6),
		},
	}

	for _, c := range cases {
		result := c.tupleA.Sub(*c.tupleB)

		if !c.result.Equal(*result) {
			t.Errorf("%s - failed adding tuple a to tuple b: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Negate(t *testing.T) {
	cases := []struct {
		name   string
		tuple  *Tuple
		result *Tuple
	}{
		{
			name:   "OK/Negating_Vector",
			tuple:  Vector(1, -2, 3),
			result: Vector(-1, 2, -3),
		},
	}

	for _, c := range cases {
		result := c.tuple.Negate()

		if !c.result.Equal(*result) {
			t.Errorf("%s - failed negating tuple: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Multiply(t *testing.T) {
	cases := []struct {
		name   string
		tuple  *Tuple
		scalar float64
		result *Tuple
	}{
		{
			name:   "OK/Multiplying_Tuple",
			tuple:  Vector(1, -2, 3),
			scalar: 3.5,
			result: Vector(3.5, -7, 10.5),
		},
	}

	for _, c := range cases {
		result := c.tuple.Multiply(c.scalar)

		if !c.result.Equal(*result) {
			t.Errorf("%s - failed multiplying tuple: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Divide(t *testing.T) {
	cases := []struct {
		name   string
		tuple  *Tuple
		scalar float64
		result *Tuple
	}{
		{
			name:   "OK/Diving_Tuple",
			tuple:  Vector(1, -2, 3),
			scalar: 2,
			result: Vector(0.5, -1, 1.5),
		},
	}

	for _, c := range cases {
		result := c.tuple.Divide(c.scalar)

		if !c.result.Equal(*result) {
			t.Errorf("%s - failed diving tuple: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Magnitude(t *testing.T) {
	cases := []struct {
		name   string
		tuple  *Tuple
		result float64
	}{
		{
			name:   "OK/Magnitude_Vector_1",
			tuple:  Vector(1, 0, 0),
			result: 1,
		},
		{
			name:   "OK/Magnitude_Vector_2",
			tuple:  Vector(0, 1, 0),
			result: 1,
		},
		{
			name:   "OK/Magnitude_Vector_3",
			tuple:  Vector(0, 0, 1),
			result: 1,
		},
		{
			name:   "OK/Magnitude_Vector_Not_unitary",
			tuple:  Vector(1, 2, 3),
			result: math.Sqrt(14),
		},
		{
			name:   "OK/Magnitude_Vector_Negative",
			tuple:  Vector(-1, -2, -3),
			result: math.Sqrt(14),
		},
		{
			name:   "The magnitude of a normalized vector",
			tuple:  Vector(1, 2, 3).Normalize(),
			result: 1,
		},
	}

	for _, c := range cases {
		result := c.tuple.Magnitude()

		if !utils.FloatEquals(result, c.result) {
			t.Errorf("%s - failed getting magnitud of tuple: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}

func TestTuple_Normalize(t *testing.T) {
	cases := []struct {
		name   string
		tuple  *Tuple
		result *Tuple
	}{
		{
			name:   "Normalizing vector(4, 0, 0) gives (1, 0, 0)",
			tuple:  Vector(4, 0, 0),
			result: Vector(1, 0, 0),
		},
		{
			name:   "Normalizing vector(1, 2, 3)",
			tuple:  Vector(1, 2, 3),
			result: Vector(0.26726, 0.53452, 0.80178),
		},
	}

	for _, c := range cases {
		result := c.tuple.Normalize()

		if !result.Equal(*c.result) {
			t.Errorf("%s - failed normalicing tuple: expected (%v) but got (%v)", c.name, c.result, result)
		}
	}
}
