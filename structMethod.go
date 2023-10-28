package main

import (
	"fmt"
)

//go is not OOP

type rectangle struct{
	Width int
	Height int
}

func main() {
	rec1 := rectangle{2,3}
	fmt.Printf("Width: %d, Height: %d\n", rec1.Width, rec1.Height)	

	fmt.Printf("Area of the rectangle: %d", rec1.calcAreaRec())


}

func (currRect rectangle) calcAreaRec() int{		//similar to method
	return currRect.Width * currRect.Height
}