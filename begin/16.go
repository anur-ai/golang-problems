package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2 float64

	fmt.Print("Enter x1 and x2 ")
	fmt.Scanln(&x1, &x2)
	fmt.Println(math.Abs(x1 - x2))
}