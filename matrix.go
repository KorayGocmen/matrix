package matrix

import (
	"errors"
)

// Matrix is the main matrix data structure
type Matrix struct {
	rowCount int
	colCount int
	matrix   [][]float64
}

// NewMatrix creates a new matrix with the given row, col and initial values.
func NewMatrix(rowCount, colCount int, values []float64) (*Matrix, error) {
	if values != nil && len(values) != rowCount*colCount {
		return nil, errors.New(MatrixInitialValuesDimensionsMismatch)
	}

	matrix := make([][]float64, rowCount)
	for i := range matrix {
		matrix[i] = make([]float64, colCount)
	}

	index := 0
	for i := range matrix {
		for j := range matrix[i] {
			ijValue := 0.0
			if values != nil {
				ijValue = values[index]
			}
			matrix[i][j] = ijValue
			index++
		}
	}

	m := Matrix{rowCount: rowCount, colCount: colCount, matrix: matrix}
	return &m, nil
}

// RowAt returns the row with the given row index
func (m Matrix) RowAt(rowIndex int) ([]float64, error) {
	if rowIndex >= m.rowCount {
		return nil, errors.New(MatrixRowIndexOutOfRange)
	}
	return m.matrix[rowIndex], nil
}

// ColAt returns the column with the given column index
func (m Matrix) ColAt(colIndex int) ([]float64, error) {
	var col []float64
	for i := 0; i < m.rowCount; i++ {
		if colIndex >= m.colCount {
			return nil, errors.New(MatrixColIndexOutOfRange)
		}
		col = append(col, m.matrix[i][colIndex])
	}
	return col, nil
}

// Dot does matrix multiplication and assigns the result to target matrix
func Dot(x, y *Matrix) (*Matrix, error) {
	if x.colCount != y.rowCount {
		return nil, errors.New(MatrixDotImpossibleDueToDimensions)
	}

	out, _ := NewMatrix(x.rowCount, y.colCount, nil)

	for rowID := 0; rowID < x.rowCount; rowID++ {
		row, _ := x.RowAt(rowID)
		for colID := 0; colID < y.colCount; colID++ {
			col, _ := y.ColAt(colID)
			for i := 0; i < len(row); i++ {
				out.matrix[rowID][colID] += row[i] * col[i]
			}
		}
	}

	return out, nil
}

// Scale performs scalar multiplication with the given factor
func Scale(A float64, x *Matrix) (*Matrix, error) {
	out, _ := NewMatrix(x.rowCount, x.colCount, nil)

	for rowID := 0; rowID < x.rowCount; rowID++ {
		for colID := 0; colID < x.colCount; colID++ {
			out.matrix[rowID][colID] = A * x.matrix[rowID][colID]
		}
	}

	return out, nil
}

// Transpose creates a new transpose matrix for the given matrix
func Transpose(x *Matrix) (*Matrix, error) {
	var outValues []float64
	for colID := 0; colID < x.colCount; colID++ {
		col, _ := x.ColAt(colID)
		for _, val := range col {
			outValues = append(outValues, val)
		}
	}

	out, _ := NewMatrix(x.colCount, x.rowCount, outValues)
	return out, nil
}

// DeepEqual checks if two given matrixes are exactly the same
func DeepEqual(x, y *Matrix) bool {
	if x.rowCount != y.rowCount || x.colCount != y.colCount {
		return false
	}

	for rowID := 0; rowID < x.rowCount; rowID++ {
		for colID := 0; colID < x.colCount; colID++ {
			if x.matrix[rowID][colID] != y.matrix[rowID][colID] {
				return false
			}
		}
	}

	return true
}
