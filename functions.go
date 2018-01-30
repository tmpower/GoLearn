package main

import (
	"fmt"
)

func plusFunc(a int, b int) int {
	return a + b
}

func plusPlusFunc(a, b, c int) int {
	return a + b + c
}

func multAndDiv(a int, b int) (int, float64) {
	return a * b, float64(a) / float64(b)
}

func sumUp(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Print(total, "\n")
}

func fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a-1)
}

func main() {
	res := plusFunc(1, 2)
	fmt.Println(res)

	res = plusPlusFunc(1, 2, 3)
	fmt.Println(res)

	m, d := multAndDiv(3, 6)
	fmt.Println(m, d)

	sumUp(1, 2)
	sumUp(1, 4, 23, 4)

	nums := []int{1, 2, 3, 4}
	sumUp(nums...)

	fmt.Println("7!", fact(7))
}
