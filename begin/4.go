package main

import (
	"fmt"
)

func main() {
	var d float32

	fmt.Println("Enter d")
	fmt.Scanln(&d)

	fmt.Println(3.14 * d)
}