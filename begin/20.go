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

	fmt.Println(float32(math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))))

}