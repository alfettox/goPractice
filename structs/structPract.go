package main

import "fmt"

func main() {
	giulietta := car{
		maker:   "Alfa Romeo",
		power:   185.5,
		country: "Italy",
		doors:   4,
	}
	
	fmt.Println("Car Maker:", giulietta.maker)
	fmt.Println("Horsepower:", giulietta.power)
	fmt.Println("Country:", giulietta.country)
	fmt.Println("Number of Doors:", giulietta.doors)
	fmt.Println("____________________")
	
	corsa := car{
		"Opel",
		75.5,
		"Unknown",
		5,
	}
	

	fmt.Println("Maker: " + corsa.maker)
	fmt.Println("Power: ", corsa.power)
	fmt.Println("Country: " + corsa.country)
	fmt.Println("Doors: ", corsa.doors)
	fmt.Println("____________________")
	

}

type car struct {
	maker   string
	power   float64
	country string
	doors   int
}
