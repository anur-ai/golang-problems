package main

import (
	"fmt"
)

func main() {
	var a, c, b int

	fmt.Print("Enter A, C, B ")
	fmt.Scanln(&a, &c, &b)

	ac := c - a
	bc := b - c

	fmt.Println(ac, bc)
}