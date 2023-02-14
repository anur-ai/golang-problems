package main

import (
	"fmt"
	"math"
)

func hype(num1 float64, num2 float64) (float64) {
	var result float64

	result = math.Sqrt(math.Pow(num1, 2) + math.Pow(num2, 2))

	return result
}

func main() {
	var a, b, c float64

	fmt.Println("Enter a, b")
	fmt.Scanln(&a, &b)

	c = hype(a, b)
	p := a + b + c

	fmt.Printf("%g, %g", c, p)
}