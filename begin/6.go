package main

import "fmt"

func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	V := a * b * c
	S := 2 * (a * b + b * c + a * c)
	fmt.Println("V = ", V)
	fmt.Println("S = ", S)
}