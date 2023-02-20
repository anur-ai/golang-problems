package main

import (
	"fmt"
	"math"
)

func main(){
	var x1, x2 float64
	fmt.Scan(&x1, &x2)
	fmt.Println(math.Abs(x2 - x1))
}