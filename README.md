# matrix

Package matrix implements a simple library for matrix operations. A matrix struct is used to create new matrixes and operations are performed using this structure.

100% Test coverage

---
#### Full Documentation:

https://godoc.org/github.com/KorayGocmen/matrix

```go
package main

import (
	"fmt"

	"github.com/koraygocmen/matrix"
)

func main() {

	x, _ := matrix.New(2, 2, []float64{
		1, 2,
		3, 4,
	})

	y, _ := matrix.New(2, 2, []float64{
		5, 6,
		7, 8,
	})

	xPlusY, _ := matrix.Add(x, y)
	fmt.Println(xPlusY)

	xMinusY, _ := matrix.Subtract(x, y)
	fmt.Println(xMinusY)

	scaledX, _ := matrix.Scale(2, x)
	fmt.Println(scaledX)

	transposedX, _ := matrix.Transpose(x)
	fmt.Println(transposedX)

	dotProduct, _ := matrix.Dot(x, y)
	fmt.Println(dotProduct)

	scalarAdded, _ := matrix.AddScalar(2, x)
	fmt.Println(scalarAdded)

	scalarSubtracted, _ := matrix.SubtractScalar(1, x)
	fmt.Println(scalarSubtracted)
}
```

---

### License

Released under the [MIT License](https://github.com/KorayGocmen/matrix/blob/master/LICENSE).