package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter A, B ")
	fmt.Scanln(&a, &b)

	fmt.Println(a - (int(a / b) * b))
}