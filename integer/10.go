package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	last_digit := num % 10
	second_digit := (num % 100 / 10)

	fmt.Println(last_digit)
	fmt.Println(second_digit)
}