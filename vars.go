package main

import (
	"fmt"
)

func main() {
	x :=5
	x =2
	y :=x		//PASS BY VALUE
	y =9
	
    output := fmt.Sprintf("x:%d, y:%d", x, y)
	fmt.Println(output)
}