package main

import (
    "fmt"
    "strings"
)

//pass a pointer type, modify the original value

func isUpper(x *string) bool {
    if strings.ToUpper(*x) == *x {
        return true
    }
    return false
}

func main() {
    var msg string = "Good morning"
    msgPtr := &msg

    fmt.Println(isUpper(msgPtr))
    fmt.Println(isLower(msgPtr))

	makeItUppercase(msgPtr)	//modify the memory value!
	fmt.Println(msg)
}

func makeItUppercase(x *string) {
	*x =strings.ToUpper(*x)
}

func isLower(y *string) bool {		
	return(strings.ToLower(*y) == *y)
}


