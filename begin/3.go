package main

import "fmt"

func main(){
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("P = ", (a + b) * 2)
	fmt.Println("S = ", a * b)
}