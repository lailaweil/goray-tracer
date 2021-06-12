package models

import (
	"github.com/laiweil/goray-tracer/utils"
	"testing"
)

func TestCreateMatrix_4x4(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "Constructing and inspecting a 4x4 matrix 1",
			a:    0,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 2",
			a:    0,
			b:    3,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 3",
			a:    1,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 4",
			a:    1,
			b:    2,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 5",
			a:    2,
			b:    2,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 6",
			a:    3,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 4x4 matrix 7",
			a:    3,
			b:    2,
		},
	}

	matrix := [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}

	result := CreateMatrix(4, 4)
	result[0][0] = 1
	result[0][3] = 4
	result[1][0] = 5.5
	result[1][2] = 7.5
	result[2][2] = 11
	result[3][0] = 13.5
	result[3][2] = 15.5

	for _, c := range cases {
		if !utils.FloatEquals(result[c.a][c.b], matrix[c.a][c.b]) {
			t.Errorf("%s - failed creating and inspecting matrix: expected to be (%v) but got (%v)", c.name, result[c.a][c.b], matrix[c.a][c.b])
		}
	}
}

func TestCreateMatrix_2x2(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "Constructing and inspecting a 2x2 matrix 1",
			a:    0,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 2x2 matrix 2",
			a:    0,
			b:    1,
		},
		{
			name: "Constructing and inspecting a 2x2 matrix 3",
			a:    1,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 2x2 matrix 4",
			a:    1,
			b:    1,
		},
	}

	matrix := [][]float64{
		{-3, 5},
		{1, -2},
	}

	result := CreateMatrix(2, 2)
	result[0][0] = -3
	result[0][1] = 5
	result[1][0] = 1
	result[1][1] = -2

	for _, c := range cases {
		if !utils.FloatEquals(result[c.a][c.b], matrix[c.a][c.b]) {
			t.Errorf("%s - failed creating and inspecting matrix: expected to be (%v) but got (%v)", c.name, result[c.a][c.b], matrix[c.a][c.b])
		}
	}
}

func TestCreateMatrix_3x3(t *testing.T) {
	cases := []struct {
		name string
		a    int
		b    int
	}{
		{
			name: "Constructing and inspecting a 3x3 matrix 1",
			a:    0,
			b:    0,
		},
		{
			name: "Constructing and inspecting a 3x3 matrix 2",
			a:    1,
			b:    1,
		},
		{
			name: "Constructing and inspecting a 3x3 matrix 3",
			a:    2,
			b:    2,
		},
	}

	matrix := [][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}

	result := CreateMatrix(3, 3)
	result[0][0] = -3
	result[1][1] = -2
	result[2][2] = 1

	for _, c := range cases {
		if !utils.FloatEquals(result[c.a][c.b], matrix[c.a][c.b]) {
			t.Errorf("%s - failed creating and inspecting matrix: expected to be (%v) but got (%v)", c.name, result[c.a][c.b], matrix[c.a][c.b])
		}
	}
}

func TestMatrix_Equal(t *testing.T) {
	cases := []struct {
		name    string
		matrixA Matrix
		matrixB Matrix
		equal   bool
	}{
		{
			name: "Identical matrices",
			matrixA: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 3},
			},
			matrixB: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 3},
			},
			equal: true,
		},
		{
			name: "Different Matrices",
			matrixA: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 3},
			},
			matrixB: [][]float64{
				{2, 3, 4, 5},
				{6, 7, 8, 9},
				{8, 7, 6, 5},
				{4, 3, 2, 1},
			},
			equal: false,
		},
	}

	for _, c := range cases {
		result := c.matrixA.Equal(c.matrixB)
		if result != c.equal {
			t.Errorf("%s - failed comparing matrices: expected to be (%t) but got (%t)", c.name, c.equal, result)
		}
	}
}

func TestMatrix_Multiply(t *testing.T) {
	cases := []struct {
		name    string
		matrixA Matrix
		matrixB Matrix
		result  Matrix
	}{
		{
			name: "Multiplying two matrices",
			matrixA: [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 8, 7, 6},
				{5, 4, 3, 2},
			},
			matrixB: [][]float64{
				{-2, 1, 2, 3},
				{3, 2, 1, -1},
				{4, 3, 6, 5},
				{1, 2, 7, 8},
			},
			result: [][]float64{
				{20, 22, 50, 48},
				{44, 54, 114, 108},
				{40, 58, 110, 102},
				{16, 26, 46, 42},
			},
		},
	}

	for _, c := range cases {
		result := c.matrixA.Multiply(c.matrixB)

		if !result.Equal(c.result) {
			t.Errorf("%s - failed multiplying matrices: expected to be (%v) but got (%v)", c.name, c.result, result)
		}
	}
}
