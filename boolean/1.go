package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Enter a: ")
	fmt.Scanln(&a)

	fmt.Println(bool(a > 0))

}