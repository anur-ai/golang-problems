package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter A: ")
	fmt.Scanln(&a)

	fmt.Println(bool(int(a / 10) <= 9 && a % 2 == 0 && int(a / 10) >= 1))
}