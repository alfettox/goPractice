package main

import (
	"fmt"
)

// essai is a higher-order function that takes an operation (a function) and two integers as arguments.
func essai(operation func(int, int) int, a, b int) int {
	result := operation(a, b)
	return result
}

func main() {
	// addition is a callback function for addition.
	addition := func(x, y int) int {
		return x + y
	}

	// Use the essai function with the addition callback.
	sum := essai(addition, 5, 3)

	fmt.Printf("The result of the addition is: %d\n", sum)

	// highFunction is another callback function for multiplication.
	highFunction := func(z, l int) int {
		return z * l
	}

	// Use the anotherFunction to apply the highFunction callback.
	res := anotherFunction(highFunction, 1, 2)

	fmt.Println("Result highFunction = ", res)
}

// anotherFunction is a higher-order function that takes someFunc (a function) and two integers as arguments.
func anotherFunction(someFunc func(int, int) int, a, b int) int {
	result := someFunc(a, b)
	return result
}
