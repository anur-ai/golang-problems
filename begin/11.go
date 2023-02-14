package main

import (
	"fmt"
)

func module(int1 float64) float64 {
	var result float64
	if (int1 > 0) {
		result = int1
	} else {
		result = int1 * (-1)
	}
	return result
}

func main() {
	var a, b float64

	fmt.Println("Enter a, b")
	fmt.Scanln(&a, &b)
	fmt.Println(a + b, a - b, a * b,module(a), module(b))
}