package main

import (
	"fmt"
	"math"
)

func main(){
	var a float64
	fmt.Scan(&a)
	fmt.Println("V = ", math.Pow(a, 3))
	fmt.Println("S = ", 6 * math.Pow(a, 2))
}