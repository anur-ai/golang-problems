package main

import (
	"fmt"
	"math"
)

func main(){
	var x1, y1, x2, y2 float64
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	fmt.Println(math.Sqrt(math.Pow((x2 - x1), 2) - math.Pow((y2 - y1), 2)))
}