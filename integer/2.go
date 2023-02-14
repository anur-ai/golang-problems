package main

import "fmt"

func main() {
	var m int

	fmt.Print("Enter M ")
	fmt.Scanln(&m)

	tons := float32(m) / 1000

	fmt.Println(tons, "tons")
}