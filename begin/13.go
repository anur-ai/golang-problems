package main 

import (
	"fmt"
)

func area(radius float64) (float64) {
	p := 3.14
	return radius * radius * p
}

func main() {
	var r1, r2 float64

	fmt.Println("Enter radius1 and radius2 ")
	fmt.Scanln(&r1, &r2)

	fmt.Print("S1=", area(r1), "S2=", area(r2), float32(area(r1) - area(r2)))
}