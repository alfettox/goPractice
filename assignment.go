package main

import (
	"fmt"
)

func main() {
	sendsSoFar := 250
	const sendsToAdd = 25
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("You have already sent: ", sendsSoFar, "messages")
	
}

func incrementSends(sendsSoFar, sendsToAdd int) int{
	sendsSoFar = sendsSoFar + sendsToAdd
	return sendsSoFar
}