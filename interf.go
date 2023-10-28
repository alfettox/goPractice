package main

/*Interface = collection of signatures */

import (
	"fmt"
	"math"
)

const RoomHeight float64 = 2.7

func main() {
	rect1 := rect{2, 3, 13.25}
	square1 := square{3, 19.90}
	circle1 := circle{2.5, 45.50}

	// RectInfo := fmt.Sprintf("Rectangle area: %f, perimeter %f", rect1.Width, rect1.Width2)
	// SquareInfo := fmt.Sprintf("Square area: %f, perimeter %f", square1.Side, square1.Side)

	printShapeInfo(rect1)
	printShapeInfo(square1)
	printShapeInfo(circle1)
}

type circle struct{
	Radius float64
	ElemCost float64
}

type shape interface {
	Area() float64
	Perimeter() float64
	Cost() float64
	Volume() float64
}

type rect struct {
	Width  float64
	Width2 float64
	ElemCost float64
}

type square struct {
	Side float64
	ElemCost float64
}

func (r rect) Area() float64 {
	return r.Width * r.Width2
}

func (r rect) Perimeter() float64 {
	return r.Width * r.Width2 * 2
}

func (s square) Area() float64 {
	return s.Side * s.Side
}

func (s square) Perimeter() float64 {
	return s.Side * 4
}

func printShapeInfo(s shape) {

	area := fmt.Sprintf("%.2f", s.Area())
	perimeter := fmt.Sprintf("%.2f", s.Perimeter())
	cost := fmt.Sprintf("%.2f", s.Cost())
	volume := fmt.Sprintf("%.2f", s.Volume())


	fmt.Printf("Area: %sm, Perimeter: %sm\n", area, perimeter)
	fmt.Printf("The cost of this element is: %s$\n", cost)
	fmt.Printf("The room volume (standard height 2.7m) is: %smc\n", volume)
	printSeparator()
}

func printSeparator(){
	fmt.Println("___________________________________")
}

func (r rect) Cost() float64 {
    return r.Area() * r.ElemCost
}

func (s square) Cost() float64 {
    return s.Area() * s.ElemCost
}

func (c circle) Area() float64{
	return math.Pi * c.Radius * c.Radius
}

func (c circle) Perimeter() float64{
	return 2 * math.Pi * c.Radius
}

func (c circle) Cost() float64{
	return c.Area() * c.ElemCost
}

func (c circle) Volume() float64{
	return c.Area()*RoomHeight
}

func (r rect) Volume() float64{
	return r.Area()*RoomHeight
}

func (s square) Volume() float64{
	return s.Area()*RoomHeight
}