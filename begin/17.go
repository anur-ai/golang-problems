package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	ac := c - a
	bc := c - b 
	fmt.Print("AC=", ac, " BC=", bc)
}