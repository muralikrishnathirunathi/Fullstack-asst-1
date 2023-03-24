package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = true
	MaxInt uint64     = 1<<63 - 1
	z      complex128 = cmplx.Sqrt(-7 + 24i)
)

func main() {
	fmt.Println("Value of ToBe:", ToBe)
	fmt.Println("Value of MaxInt:", MaxInt)
	fmt.Println("Value of z:", z)
}
