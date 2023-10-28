package main

// Nested structs, struct within another struct,
// but there's no direct relationship between the inner and outer structs.
// Each struct retains its own identity and fields.
// no strong "has-a" or "is-a" relationship

import (
    "fmt"
    "time"
)

func main() {
    subject1Info := Sender{User{"Marie", 25}, parseTime("22:45")}
    subject2Info := Receiver{User{"Paulie", 28}, parseTime("22:48")}
    fmt.Println(subject1Info)
    fmt.Println(subject2Info)

    // fmt.Println(subject1Info.Age)  not at the top level, not working as embedded
}

type User struct {
    Name string
    Age  int
}

type Sender struct {
    User      User
    SentTime  time.Time
}

type Receiver struct {
    User         User
    ReceivedTime time.Time
}

func parseTime(timeStr string) time.Time {
    layout := "15:04"
    t, err := time.Parse(layout, timeStr)
    if err != nil {
        panic(err)
    }
    return t
}

func (s Sender) String() string {		//Override satisfying stringer interface
    return fmt.Sprintf("Name: %s, Age: %d, SentTime: %s", s.User.Name, s.User.Age, s.SentTime.Format("15:04:05"))
}

func (s Receiver) String() string {		//Override satisfying stringer interface
    return fmt.Sprintf("Name: %s, Age: %d, ReceivedTime: %s", s.User.Name, s.User.Age, s.ReceivedTime.Format("15:04:05"))
}
