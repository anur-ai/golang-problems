package main

import "fmt"

func main(){
	var a, b ,c int
	fmt.Scan(&a, &c, &b)
	AC := c - a
	BC := b - c
	fmt.Println(AC * BC)
}