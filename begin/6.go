package main

import (
	"fmt"
)

func main() {
	var a, b, c, v, s float64

	fmt.Println("Enter a, b, c")
	fmt.Scanln(&a, &b, &c)

	v = a * b * c
	s = 2 * (a * b + b * c + a * c)

	fmt.Print("V=", v, " S=", s)
}