package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Enter a, b")
	fmt.Scanln(&a, &b)

	fmt.Println(a + b, a - b, a * b, math.Pow(a, 2), math.Pow(b, 2))
}