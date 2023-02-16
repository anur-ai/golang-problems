package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter A: ")
	fmt.Scanln(&a)

	fmt.Println(bool(int(a / 100) <= 9 && a % 2 == 1 && int(a / 100) >= 1))
}