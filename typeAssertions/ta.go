package main

import (
	"fmt"
)

func main() {
	test(email{
		isSubscribed: true,
		body:        "Ciao mamma",
		destination: "Rue de Montreal 22",
	})
	test(email{
		isSubscribed: false,
		body:         "Goodmoring Seattle",
		destination:  "Rue de Narbonne 33",
	})
	test(email{
		isSubscribed: false,
		body:         "Hi All!",
		destination:  "Rue Vancouver 1",
	})
	test(email{
		isSubscribed: true,
		body:         "gg!",
		destination:  "Rue Norvège 2232",
	})
	test(sms{
		isSubscribed:  true,
		body:         "Hello",
		toPhoneNumber: "1234567890",
	})
	test(sms{
		isSubscribed:  false,
		body:         "Greetings",
		toPhoneNumber: "9876543210",
	})
	test(invalid{})
}

func getReport(exp expense) (string, float64) {
	switch e := exp.(type) {
	case email:
		return e.destination, e.cost()
	case sms:
		return e.toPhoneNumber, e.cost()
	default:
		return "", 0.0
	}
}

func (currEmail email) cost() float64 {
	if !currEmail.isSubscribed {
		return float64(len(currEmail.body)) * 0.05
	} else {
		return float64(len(currEmail.body)) * 0.01
	}
}

func (currSms sms) cost() float64 {
	if !currSms.isSubscribed {
		return float64(len(currSms.body)) * 0.1
	}
	return float64(len(currSms.body))
}

func (inv invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	destination  string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct {
}

func estYearCost(exp expense, avgMsgPerYear int) float64 {
	return exp.cost() * float64(avgMsgPerYear)
}

func test(exp expense) {
	address, cost := getReport(exp)
	switch exp.(type) {
	case email:
		fmt.Printf("The email to %v cost %.2f\n", address, cost)
	case sms:
		fmt.Printf("The sms to %v cost %.2f\n", address, cost)
	default:
		fmt.Println("Invalid")
	}
	fmt.Println("---------------------------")
}
