package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter A, B, C: ")
	fmt.Scanln(&a,&b,&c)

	fmt.Println(bool(a < b &&  b < c || a > b && b > c))
}