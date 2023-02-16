package main

import "fmt"

func main() {
	var a,b,c int

	fmt.Print("Enter A, B, C: ")
	fmt.Scanln(&a,&b,&c)

	var num = [3]int{a, b, c}

	positive := 0
	negative := 0

	for i := 0; i < 3; i++ {
		if (num[i] > 0) {
			positive += 1
		} else {
			negative += 1
		}
	}

	fmt.Println("Positives:", positive, "Negatives:", negative)
}