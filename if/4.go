package main

import "fmt"

func main() {
	var a,b,c int

	fmt.Print("Enter A, B, C: ")
	fmt.Scanln(&a,&b,&c)

	var num = [3]int{a, b, c}

	sum := 0

	for i := 0; i < 3; i++ {
		if (num[i] > 0) {
			sum += 1
		}
	}

	fmt.Println(sum)
}