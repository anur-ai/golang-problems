package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter A ")
	fmt.Scanln(&a)

	fmt.Println(int(a / 10))
	fmt.Println(a - (int(a / 10) * 10))
}