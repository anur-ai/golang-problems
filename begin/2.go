package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64

	fmt.Println("Enter a:")
	fmt.Scanln(&a)
	fmt.Println(math.Pow(a, 2))
}