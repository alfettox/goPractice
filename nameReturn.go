package main

import (
    "fmt"
)

func main() {
    x, y := getCoordinates2()
    fmt.Println(x, y)

    resultX, _ := getCoordinates1()

    // fmt.Println(result)
    fmt.Println(resultX)
}

func getCoordinates1() (x int, y int) {
    x = 4
    y = 2
    return
}

func getCoordinates2() (int, int) {
    var x, y int
    x = 4
    y = 2
    return x, y
}
