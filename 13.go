package main

import (
	"fmt"
	"math"
)

func main() {
	x, y := 3.0, 4.0
	hypotenuse := math.Sqrt(x*x + y*y)
	fmt.Printf("The hypotenuse of a right triangle with sides %.2f and %.2f is %.2f.\n", x, y, hypotenuse)
}
