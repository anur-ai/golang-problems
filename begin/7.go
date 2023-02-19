package main

import (
	"fmt"
	"math"
)

func main(){
	var R float64
	P := 3.14
	fmt.Scan(&R)
	L := 2 * P * R
	S := P * math.Pow(R, 2)
	fmt.Println("L = ", L)
	fmt.Println("S = ", S)
}