package models

import "github.com/laiweil/goray-tracer/utils"

type Matrix [][]float64

func CreateMatrix(rows int, columns int) Matrix {
	matrix := make([][]float64, rows)
	matrixRows := make([]float64, columns*rows)
	for i := 0; i < rows; i++ {
		matrix[i] = matrixRows[i*columns : (i+1)*columns]
	}

	return matrix
}

func (m Matrix) Equal(matrixB Matrix) bool {
	for indexRow, row := range m {
		for indexElement, element := range row {
			if !utils.FloatEquals(element, matrixB[indexRow][indexElement]) {
				return false
			}
		}
	}
	return true
}

func (m Matrix) Multiply(matrixB Matrix) Matrix {
	length := len(m)
	result := CreateMatrix(length, length)

	for row := 0; row < length; row++ {
		for col := 0; col < length; col++ {
			for i := 0; i < length; i++ {
				result[row][col] = result[row][col] + (m[row][i] * matrixB[i][col])
			}
		}
	}
	return result
}
