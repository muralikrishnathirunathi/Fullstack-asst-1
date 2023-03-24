package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num1 := rand.Intn(100)
	num2 := rand.Intn(100)
	
	fmt.Println("First random number:", num1)
	fmt.Println("Second random number:", num2)
	
	if num1 > num2 {
		fmt.Println("The greater number is:", num1)
	} else {
		fmt.Println("The greater number is:", num2)
	}
}
