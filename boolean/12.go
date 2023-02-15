package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter A, B, C: ")
	fmt.Scanln(&a,&b, &c)

	fmt.Println(bool(a > 0 && b > 0 && c > 0))
}