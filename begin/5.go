package main

import (
	"fmt"
	"math"
)

func main() {
	var a, s, v float64

	fmt.Println("Enter a")
	fmt.Scanln(&a)

	s = 6 * math.Pow(a, 2)
	v = math.Pow(a, 3)

	fmt.Print("S=", s, " V=", v)
}