package main

import "fmt"

func main() {
	var a, b float64

	fmt.Println("Enter a, b")
	fmt.Scanln(&a, &b)

	s := (a + b) / 2

	fmt.Println(s)
}