package main

import (
	"fmt"
)

func main() {
	var l, s, r float64

	fmt.Println("Enter R")
	fmt.Scanln(&r)

	l = 2 * 3.14 * r
	s = 3.14 * r * r

	fmt.Print("L=", l, " S=", s)
}