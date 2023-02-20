package main

import (
	"fmt"
	"math"
)

func main()  {
	var a, b float64
	fmt.Scan(&a, &b)
	if a > 0 && b > 0{
		abs_a := math.Abs(a)
		abs_b := math.Abs(b)
		fmt.Println(abs_a + abs_b, abs_a - abs_b, abs_a * abs_b, abs_a / abs_b)
	}
}