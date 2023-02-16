package main

import "fmt"

func main() {
	var a,b int

	fmt.Print("Enter A, B: ")
	fmt.Scanln(&a,&b)

	if (a < b) {
		fmt.Println(a, b)
	} else {
		fmt.Println(b, a)
	}
	
}