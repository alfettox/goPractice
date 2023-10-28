package main

import (
	"fmt"
	"strconv"
)

func main() {
    x := 5
    fmt.Println("Value before the modification: " + strconv.Itoa(x))

    // pass by value
    modifyValue(x)
    fmt.Println("Var x after modification:", x) // Pass by value

    // pass by reference
    modifyValueRef(&x)		// 1 PASS THE POINTER
    fmt.Println("Var x after modification using modifyValueRef:", x)
}

func modifyValue(value int) {
    value = 10
    fmt.Println("Value inside the function (modifyValue):", value)
}

func modifyValueRef(value *int) { // 2USE POINTER TYPE
    *value = 55					  // 3 USE POINTER REFERENCE
    fmt.Println("Value inside the function (modifyValueRef):", *value)
}
