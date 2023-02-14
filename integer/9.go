package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Enter a: ")
	fmt.Scanln(&a)

	a = int(a / 100)
	fmt.Println(a)
}