package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64

	fmt.Print("Enter x1 and y1 ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("Enter x2 and y2 ")
	fmt.Scanln(&x2, &y2)

	height := math.Abs(y2 - y1)
	width := math.Abs(x2 - x1)

	area := height * width

	fmt.Println(area)

}