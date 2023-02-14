package main

import "fmt"

func main() {
	var l int

	fmt.Print("Enter L ")
	fmt.Scanln(&l)

	meter := float32(l) / 100

	fmt.Println(meter, "meters")
}