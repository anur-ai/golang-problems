package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter A, B: ")
	fmt.Scanln(&a,&b)

	fmt.Println(bool(a % 2 == 1 && b % 2 == 1))
}