package main

import "fmt"

func yearsUntil(age int) (int, int, int) {
	adult := 18 - age
	if adult < 0 {
		adult = 0
	}
	drink := 21 - age
	if drink < 0 {
		drink = 0
	}
	car := 25 - age
	if car < 0 {
		car = 0
	}
	return adult, drink, car
}

func test(age int) {
	adult, drink, car := yearsUntil(age)
	fmt.Println(adult, drink, car, "\n")
	fmt.Println("Age:", age)
	fmt.Printf("Adulthood in %d years\n", adult)
	fmt.Printf("Legal drinking age in %d years\n", drink)
	fmt.Printf("Car rental age in %d years\n", car)
	fmt.Println()
}

func main() {
	test(4)
	test(10)
	test(22)
	test(35)
}
