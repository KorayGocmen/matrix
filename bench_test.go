package matrix

import (
	"math/rand"
	"testing"
)

func BenchmarkDotProduct(b *testing.B) {

	firstMatrixValues, secondMatrixValues := []float64{}, []float64{}
	squareMatrixWidth := 1000
	for i := 0; i < squareMatrixWidth; i++ {
		for j := 0; j < squareMatrixWidth; j++ {
			firstMatrixValues = append(firstMatrixValues, rand.Float64())
			secondMatrixValues = append(secondMatrixValues, rand.Float64())
		}
	}

	x1, _ := NewMatrix(squareMatrixWidth, squareMatrixWidth, firstMatrixValues)
	x2, _ := NewMatrix(squareMatrixWidth, squareMatrixWidth, secondMatrixValues)

	Dot(x1, x2)
}

func BenchmarkScale(b *testing.B) {

	matrixValues := []float64{}
	squareMatrixWidth := 8000
	for i := 0; i < squareMatrixWidth; i++ {
		for j := 0; j < squareMatrixWidth; j++ {
			matrixValues = append(matrixValues, rand.Float64())
		}
	}

	x, _ := NewMatrix(squareMatrixWidth, squareMatrixWidth, matrixValues)

	Scale(rand.Float64(), x)
}
