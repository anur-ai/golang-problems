package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Println("Enter a")
	fmt.Scanln(&a)

	var p = a * 4
	fmt.Println("Result:", p)
}