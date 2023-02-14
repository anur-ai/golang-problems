package main

import (
	"fmt"
	"math"
)

func main() {
	var s, r float64
	p := 3.14

	fmt.Print("Enter L: ")
	fmt.Scanln(&s)

	r = math.Sqrt(s / p)
	l := r * 2 * p

	fmt.Print("D=", float32(r * 2), " L=", float32(l))
}