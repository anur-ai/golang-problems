package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Enter A: ")
	fmt.Scanln(&a)

	fmt.Println(bool(a % 2 == 1))
}