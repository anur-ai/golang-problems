package main

import "fmt"

func main() {
	var f int

	fmt.Print("Enter File size ")
	fmt.Scanln(&f)

	kb := float32(f) / 1024 // Conversion of byte into kilobyte

	if (kb == 1) {
		fmt.Println(kb, "kilobyte")
	} else {
		fmt.Println(kb, "kilobytes")
	}
}