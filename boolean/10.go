package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter A, B: ")
	fmt.Scanln(&a,&b)

	if (a % 2 == 1) {
		fmt.Println("true")
	} else if (b % 2 == 1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}