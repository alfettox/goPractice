package main

import (
	"fmt"
)

func main() {
	firstName, _ := getNames()		//ignore the last name return
	
	fmt.Println("Hello ", firstName)
}

func getNames()(string,string){
	return "Giovanni","De Franceschi"
}