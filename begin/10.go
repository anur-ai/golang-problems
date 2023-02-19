package main

import (
	"fmt"
	"math"
)

func main(){
	var a, b float64
	fmt.Scan(&a, &b)
	if a > 0 && b > 0{
		fmt.Println(math.Pow(a, 2) - math.Pow(b, 2), math.Pow(a, 2) + math.Pow(b, 2), math.Pow(a, 2) * math.Pow(b, 2), math.Pow(a, 2) / math.Pow(b, 2))
	}
} 