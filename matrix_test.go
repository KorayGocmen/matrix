package matrix

import (
	"testing"
)

func TestSuccessfulCreateMatrix(t *testing.T) {
	rowCount1 := 2
	colCount1 := 2
	x1, _ := NewMatrix(rowCount1, colCount1, []float64{
		4, -1,
		0, 5,
	})
	if x1.rowCount != rowCount1 || x1.colCount != colCount1 {
		t.Errorf("row or col count was wrong")
	}

	rowCount2 := 2
	colCount2 := 1
	x2, _ := NewMatrix(rowCount2, colCount2, []float64{
		1,
		2,
	})
	if x2.rowCount != rowCount2 || x2.colCount != colCount2 {
		t.Errorf("row or col count was wrong")
	}

	rowCount3 := 1
	colCount3 := 2
	x3, _ := NewMatrix(rowCount3, colCount3, []float64{
		1, 2,
	})
	if x3.rowCount != rowCount3 || x3.colCount != colCount3 {
		t.Errorf("row or col count was wrong")
	}
}

func TestUnsuccessfulCreateMatrix(t *testing.T) {
	rowCount1 := 2
	colCount1 := 3
	_, err1 := NewMatrix(rowCount1, colCount1, []float64{
		4, -1,
		0, 5,
	})
	if err1 == nil || err1.Error() != MatrixInitialValuesDimensionsMismatch {
		t.Errorf("row or col count was wrong")
	}

	rowCount2 := 1
	colCount2 := 2
	_, err2 := NewMatrix(rowCount2, colCount2, []float64{
		4, 2,
		0,
	})
	if err2 == nil || err2.Error() != MatrixInitialValuesDimensionsMismatch {
		t.Errorf("row or col count was wrong")
	}
}

func TestSuccessfulRowAt(t *testing.T) {
	x1, _ := NewMatrix(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x2, _ := NewMatrix(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})

	x1RowAtReal, _ := x1.RowAt(0)
	x1RowAtExpected := []float64{4, -1}
	if !arrayEqual(x1RowAtReal, x1RowAtExpected) {
		t.Errorf("row at error, %v, %v", x1RowAtReal, x1RowAtExpected)
	}

	x2RowAtReal, _ := x2.RowAt(0)
	x2RowAtExpected := []float64{1, 8, 0}
	if !arrayEqual(x2RowAtReal, x2RowAtExpected) {
		t.Errorf("row at error, %v, %v", x2RowAtReal, x2RowAtExpected)
	}

	x3RowAtReal, _ := x2.RowAt(1)
	x3RowAtExpected := []float64{6, -2, 3}
	if !arrayEqual(x3RowAtReal, x3RowAtExpected) {
		t.Errorf("row at error, %v, %v", x3RowAtReal, x3RowAtExpected)
	}
}

func TestUnsuccessfulRowAt(t *testing.T) {
	x1, _ := NewMatrix(2, 2, []float64{
		4, -1,
		0, 5,
	})

	_, err1 := x1.RowAt(2)
	if err1.Error() != MatrixRowIndexOutOfRange {
		t.Errorf("row at error")
	}
}

func TestSuccessfulColAt(t *testing.T) {
	x1, _ := NewMatrix(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x2, _ := NewMatrix(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})

	x1ColAtReal, _ := x1.ColAt(0)
	x1ColAtExpected := []float64{4, 0}
	if !arrayEqual(x1ColAtReal, x1ColAtExpected) {
		t.Errorf("col at error, %v, %v", x1ColAtReal, x1ColAtExpected)
	}

	x2ColAtReal, _ := x2.ColAt(0)
	x2ColAtExpected := []float64{1, 6}
	if !arrayEqual(x2ColAtReal, x2ColAtExpected) {
		t.Errorf("col at error, %v, %v", x2ColAtReal, x2ColAtExpected)
	}

	x3ColAtReal, _ := x2.ColAt(1)
	x3ColAtExpected := []float64{8, -2}
	if !arrayEqual(x3ColAtReal, x3ColAtExpected) {
		t.Errorf("col at error, %v, %v", x3ColAtReal, x3ColAtExpected)
	}
}

func TestUnsuccessfulColAt(t *testing.T) {
	x1, _ := NewMatrix(2, 2, []float64{
		4, -1,
		0, 5,
	})

	_, err1 := x1.ColAt(2)
	if err1.Error() != MatrixColIndexOutOfRange {
		t.Errorf("row at error")
	}
}

func TestSuccessfulDotProduct(t *testing.T) {
	x1, _ := NewMatrix(2, 3, []float64{
		1, 2, 3,
		4, 5, 6,
	})
	x2, _ := NewMatrix(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})
	expectedResult1, _ := NewMatrix(2, 2, []float64{
		58, 64,
		139, 154,
	})

	result1, err1 := Dot(x1, x2)

	if err1 != nil || !DeepEqual(expectedResult1, result1) {
		t.Errorf("dot product error")
	}

	x3, _ := NewMatrix(2, 2, []float64{
		4, -1,
		0, 5,
	})
	x4, _ := NewMatrix(2, 3, []float64{
		1, 8, 0,
		6, -2, 3,
	})
	expectedResult2, _ := NewMatrix(2, 3, []float64{
		-2, 34, -3,
		30, -10, 15,
	})

	result2, err2 := Dot(x3, x4)

	if err2 != nil || !DeepEqual(expectedResult2, result2) {
		t.Errorf("dot product error")
	}
}

func TestUnsuccessfulDotProduct(t *testing.T) {
	x1, _ := NewMatrix(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})
	x2, _ := NewMatrix(3, 2, []float64{
		7, 8,
		9, 10,
		11, 12,
	})

	_, err1 := Dot(x1, x2)

	if err1.Error() != MatrixDotImpossibleDueToDimensions {
		t.Errorf("dot product error")
	}
}

func TestSuccessfulScale(t *testing.T) {
	x1, _ := NewMatrix(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	expectedResult1, _ := NewMatrix(2, 3, []float64{
		3, 6, 9,
		15, 18, 21,
	})

	result1, err1 := Scale(3, x1)

	if err1 != nil || !DeepEqual(expectedResult1, result1) {
		t.Errorf("scale error")
	}

	x2, _ := NewMatrix(2, 2, []float64{
		1, 2,
		5, 6,
	})

	expectedResult2, _ := NewMatrix(2, 2, []float64{
		-2, -4,
		-10, -12,
	})

	result2, err2 := Scale(-2, x2)

	if err2 != nil || !DeepEqual(expectedResult2, result2) {
		t.Errorf("scale error")
	}
}
func TestSuccessfulDeepEqual(t *testing.T) {
	x1, _ := NewMatrix(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	x2, _ := NewMatrix(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	if !DeepEqual(x1, x2) {
		t.Errorf("deep equal error")
	}

	x3, _ := NewMatrix(1, 1, []float64{
		1,
	})

	x4, _ := NewMatrix(1, 1, []float64{
		1,
	})

	if !DeepEqual(x3, x4) {
		t.Errorf("deep equal error")
	}
}

func TestUnsuccessfulDeepEqual(t *testing.T) {
	x1, _ := NewMatrix(2, 4, []float64{
		1, 2, 3, 4,
		5, 6, 7, 8,
	})

	x2, _ := NewMatrix(2, 3, []float64{
		1, 2, 3,
		5, 6, 7,
	})

	if DeepEqual(x1, x2) {
		t.Errorf("deep equal error")
	}

	x3, _ := NewMatrix(2, 2, []float64{
		1, 2,
		3, 4,
	})

	x4, _ := NewMatrix(2, 2, []float64{
		1, 2,
		3, 5,
	})

	if DeepEqual(x3, x4) {
		t.Errorf("deep equal error")
	}
}

func arrayEqual(arr1, arr2 []float64) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}