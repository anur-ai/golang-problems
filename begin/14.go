package main

import (
	"fmt"
	"math"
)

func main(){
	var L float64
	fmt.Scan(&L)
	P := 3.14
	R := L / (2 * P) 
	S := P * math.Pow(R, 2)
	fmt.Println("R = ", R)
	fmt.Println("S = ", S)
}