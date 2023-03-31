package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

// Function to update the age of a person using pass by value
func updateAge(p Person) {
    p.Age = 30
    fmt.Println("Inside updateAge:", p)
}

// Function to update the age of a person using pass by reference
func updateAgeByRef(p *Person) {
    p.Age = 30
    fmt.Println("Inside updateAgeByRef:", *p)
}

func main() {
    person := Person{Name: "John", Age: 25}

    // Call updateAge with pass by value
    updateAge(person)
    fmt.Println("After updateAge:", person)

    // Call updateAgeByRef with pass by reference
    updateAgeByRef(&person)
    fmt.Println("After updateAgeByRef:", person)
}
