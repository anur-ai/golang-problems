package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("Enter a and b")
	fmt.Scanln(&a, &b)

	fmt.Println(a * b)
	fmt.Println((a + b) * 2)
}