package main

import "fmt"

type email struct {
    isSubscribed bool
    body         string
}

type expense interface {
    cost() float64
}

type printer interface {
    print()
}

func (e email) cost() float64 {
    if e.isSubscribed {
        return 0.01 * float64(len(e.body))
    }
    return 0.05 * float64(len(e.body))
}

func (e email) print() {
    fmt.Printf("Email Body: %s\n", e.body)
}

func main() {
    email1 := email{isSubscribed: false, body: "Good morning, not a member"}
    test(email1)

    email2 := email{isSubscribed: true, body: "Hello member! You got a special price"}
    test(email2)

    email3 := email{isSubscribed: false, body: "This is a reminder, subscribe!"}
    test(email3)
}

func test(e email) {
    fmt.Printf("Printing cost: $%.2f\n", e.cost())
    e.print()
    separator()
}

func separator() {
    fmt.Println("-------------------------------------------------")
}
