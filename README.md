# matrix

Package matrix implements a simple library for matrix operations. A matrix struct is used to create new matrixes and operations are performed using this structure.

100% Test coverage

---
#### Full Documentation:

https://godoc.org/github.com/KorayGocmen/matrix

```go
x, _ := NewMatrix(2, 2, []float64{
  1, 2,
  3, 4,
})

y, _ := NewMatrix(2, 2, []float64{
  5, 6,
  7, 8,
})

xPlusY, _ := Add(x, y)

xMinusY, _ := Subtract(x, y)

scaledX, _ := Scale(2, x)

transposedX, _ := Transpose(x)

dotProduct, _ := Dot(x, y)

```

---

### License

Released under the [MIT License](https://github.com/KorayGocmen/matrix/blob/master/LICENSE).