package main

import (
	"fmt"
)

func main() {
	myCar := struct{
				Make string
				Model string}{ Make: "Piaggio",
							   Model: "Ape"}
	fmt.Println(myCar)
}


//Can't use with anonymous struct
// func (s myCar) string{
// 	return fmt.Sprintf("Make:",s.Make, "Model:",s.Model)
// }