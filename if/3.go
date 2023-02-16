package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter A: ")
	fmt.Scanln(&a)

	if a > 0 {
		fmt.Println(a + 1)
	} else if (a == 0) {
		fmt.Println(a + 10)
	} else {
		fmt.Println(a - 2)
	}
}