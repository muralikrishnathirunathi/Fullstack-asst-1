package main

import "fmt"

func add(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(1, 2, 3, 4, 5))
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
