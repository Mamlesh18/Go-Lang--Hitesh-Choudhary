package main

import (
	"fmt"
)

func main() {
	fmt.Println("hey i am main")
	greeting()
	x := 10
	y := 20
	fmt.Println(adder(x, y))
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(proAdder(arr...))
}

func greeting() {
	fmt.Println("hey i am greeting")

}

func adder(x int, y int) int {
	total := x + y
	return total

}

func proAdder(result ...int) int {
	total := 0
	for i := 0; i < len(result); i++ {
		total += result[i]

	}
	return total
}
