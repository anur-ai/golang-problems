package main

import (
	"fmt"
	"math"
)

func main(){
	p := 3.14
	var S float64
	fmt.Scan(&S)
	R := math.Sqrt(S * p)
	L := 2 * p *R
	D := 2 * R
	fmt.Println("D = ", D, "L = ", L)
}