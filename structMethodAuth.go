package main

import (
    "fmt"
)

type user struct {
    aD authData
}

func main() {
    currUser := user{aD: authData{"Giovanni", "1234"}}
    fmt.Println(currUser.aD.pass, currUser.aD.user)
    test(currUser.aD)
}

type authData struct {
    user string
    pass string
}

func test(auth authData) {
    fmt.Println(auth.getAuth())
}

func (a authData) getAuth() string {
    return a.user +": "+ a.pass
}
