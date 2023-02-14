package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Print("Enter A, B, C ")
	fmt.Scanln(&a, &b, &c)

	ac := c - a // Length of AC
	bc := c - b // Length of BC

	fmt.Print("AC=", ac, " BC=", bc)
}