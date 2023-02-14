package main

import (
	"fmt"
)

func changedigit(num1 int) int {
	first := int(num1 / 10)
	second := num1 - int(num1 / 10) * 10

	newnumber := second * 10 + first

	return newnumber
}