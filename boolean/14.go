package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter A, B, C: ")
	fmt.Scanln(&a,&b, &c)

	if (a > 0) {
		fmt.Println("true")
	} else if (b > 0) {
		fmt.Println("true")
	} else if (c > 0) {
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}