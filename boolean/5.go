package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Enter A and B: ")
	fmt.Scanln(&a, &b)

	fmt.Println(bool(a>=0 || b<=-2))
}