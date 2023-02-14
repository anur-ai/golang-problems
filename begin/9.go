package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Println("Enter a, b")
	fmt.Scanln(&a, &b)

	s := math.Sqrt(a * b)

	fmt.Println(s)
}