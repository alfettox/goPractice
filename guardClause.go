package main

import (
    "fmt"
    "errors"
)

func main() {
    a, errA := divide(4, 3)
    b, errB := divide(3, 0)
    fmt.Println("a:", a, "b:", b)
    fmt.Println("errA:", errA)
    fmt.Println("errB:", errB)
}

func divide(dividend, div int) (float64, error) {
    if div == 0 {
        return 0, errors.New("Divided by zero")
    }
    return float64(dividend) / float64(div), nil
}
