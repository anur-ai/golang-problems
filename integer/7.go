package main

import (
	"fmt"
)

func sum(num1 int) int {
	first := int(num1 / 10)
	second := num1 - int(num1 / 10) * 10

	return first + second
}

func subtraction(num1 int) int {
	first := int(num1 / 10)
	second := num1 - int(num1 / 10) * 10

	return first - second
}

func main() {
	var a int

	fmt.Print("Enter A: ")
	fmt.Scan(&a)

	fmt.Println(sum(a))
	fmt.Println(subtraction(a))
}