package main

import (
	"fmt"
)

func main() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	first_digit := int(num / 100)
	last_digit := num % 10
	second_digit := (num % 100 / 10)

	fmt.Println(first_digit + second_digit + last_digit)
}