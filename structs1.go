package main

import (
	"fmt"
)

func main() {
	isComplete := test(car{
		make:   "Fiat",
		doors:  5,
		height: 1.50,
		width:  1.85,
	})
	fmt.Println("The fields are complete: ", isComplete)
}

type car struct {
	make   string
	doors  int
	height float64
	width  float64
}

func test(currCar car) bool {
	complete := false
	if currCar.make != "" && currCar.doors != 0 &&currCar.height != 0.0 && currCar.width != 0.0 {
		complete = true
	}
	return complete
}
