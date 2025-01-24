package test_pkg

import "fmt"

type MyInt int

type numeric interface {
	~int | float64
}

func iGenericAdd[T numeric](a, b T) T {
	return a + b
}
func main() {
	fmt.Println(iGenericAdd(MyInt(1), MyInt(2)))
}
