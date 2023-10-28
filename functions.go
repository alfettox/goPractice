package main

import (
	"fmt"
)

func main() {
	var s1 string = "Hello"
	var s2 string = "Goodbye"

	concatenatedString := concat(s1, s2)
	fmt.Println(concatenatedString)

	sum := add2nums(1,3)
	fmt.Println(sum)

}

func concat(s1 string, s2 string) string {
	return s1 + " " + s2
}

func add2nums(n1, n2 int) int{
	sum := n1+n2

	return sum
}

func essai (func(int,int) int, int) int {

}


