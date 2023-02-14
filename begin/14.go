package main

import (
	"fmt"
	"math"
)

func main() {
	var l, r float64
	p := 3.14

	fmt.Print("Enter L: ")
	fmt.Scanln(&l)

	r = l / 2 / p

	s := p * math.Pow(r, 2)

	fmt.Print("S=", float32(s), " R=", float32(r))
}