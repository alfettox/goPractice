package main
// Fields and methods of the embedded struct are automatically promoted to the outer struct


import (
	"fmt"
)

func main() {

	newCar := cabrio{
		car: car {"Mazda", "MX-5"},
		doors: 5,
	}

	fmt.Println(newCar)

	fmt.Println(newCar.Make, newCar.Model)

}


type car struct{
	Make string
	Model string
}

type cabrio struct{   // Embedded Structs= Composition
	car
	doors int
}

