package main

import "fmt"

func split(sum int) (x, y, z int) {
	x = sum * 3 / 8
	y = sum * 2 / 8
	z = sum - x - y
	return
}

func main() {
	fmt.Println(split(24))
}
