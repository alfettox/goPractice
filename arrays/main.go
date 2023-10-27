package main

import (
	"fmt"
)

func main() {
	myIntArray := make([]int, 3)
	fmt.Println(myIntArray)

	myStringArray := make([]string, 5)
	fmt.Println(myStringArray)
	myStringArray[0] = "Giovanni"
	myStringArray[1] = "Backend Developer"
	fmt.Println(myStringArray)
	fmt.Println(len(myStringArray))

	inferredLengthArray := [...]int{1, 2, 3, 4, 5} //[...] it's a placeholder

	fmt.Println(len(inferredLengthArray))

	inferredLengthSlice := make([]int, 0) //slice

	var x int
	for x = 0; x < 2; x++ {
		inferredLengthSlice = append(inferredLengthSlice, x)
	}

	fmt.Println(inferredLengthSlice)

	newArray := [5]int{0, 1, 2, 3, 4}

	for index, value := range newArray {
		fmt.Printf("This is the key: %d, this is the value %d\n", index, value)
	}

	/*
		%d: Format as a decimal integer.
		%s: Format as a string.
		%f: Format as a floating-point number.
	*/

	anotherArray := [...]int{0, 1, 2, 3}

	for i, v := range anotherArray {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	twoDimArray := [2][3] int{
		{0,1},
		{2,3,4},
	}

	fmt.Println(twoDimArray)

	var a int = 9
	var b = &a
	fmt.Println(b)
	fmt.Println(&b)
}
