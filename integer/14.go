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

	fmt.Println(second_digit * 10 + last_digit * 100 + first_digit)
}