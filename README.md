# matrix
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

scaledX, _ := Scale(2, x)

transposedX, _ := Transpose(x)

dotProduct, _ := Dot(x, y)
```

100% Test coverage