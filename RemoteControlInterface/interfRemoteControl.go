package main

import "fmt"

type TestRemote interface {
    TurnOn()
    TurnOff()
    SetVolume(volume int)
    IncreaseVolume()
    DecreaseVolume()
    NextChannel()
    PreviousChannel()
}

type TVRemote struct {
    deviceName   string
    isTurnedOn   bool
    volume       int
    currChannel  int
    maxVolume    int
    maxChannel   int
}

type RadioRemote struct {
    deviceName   string
    isTurnedOn   bool
    volume       int
    currChannel  int
    maxVolume    int
    maxChannel   int
}

func (r *TVRemote) TurnOn() {
    r.isTurnedOn = true
    fmt.Printf("%s is turned on\n", r.deviceName)
}

func (r *TVRemote) TurnOff() {
    if r.isTurnedOn {
        r.isTurnedOn = false
        fmt.Printf("%s is off\n", r.deviceName)
    }
}

func (r *TVRemote) SetVolume(volume int) {
    if r.isTurnedOn {
        if volume >= 0 && volume <= r.maxVolume {
            r.volume = volume
            fmt.Printf("%s volume set to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s volume out of range (0 - %d)\n", r.deviceName, r.maxVolume)
        }
    } else {
        fmt.Printf("%s is off -> Can't change volume level\n", r.deviceName)
    }
}

func (r *TVRemote) IncreaseVolume() {
    if r.isTurnedOn {
        if r.volume < r.maxVolume {
            r.volume++
            fmt.Printf("%s volume increased to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s is already at max volume\n", r.deviceName)
        }
    } else {
        fmt.Printf("%s is off -> Can't increase volume.\n", r.deviceName)
    }
}

func (r *TVRemote) DecreaseVolume() {
    if r.isTurnedOn {
        if r.volume > 0 {
            r.volume--
            fmt.Printf("%s volume decreased to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s is already at min volume\n", r.deviceName)
        }
    } else {
        fmt.Printf("%s is off -> Can't change volume level\n", r.deviceName)
    }
}

func (r *TVRemote) NextChannel() {
    if r.isTurnedOn {
        r.currChannel++
        if r.currChannel > r.maxChannel {
            r.currChannel = 1
        }
        fmt.Printf("%s channel changed to %d\n", r.deviceName, r.currChannel)
    } else {
        fmt.Printf("%s is off -> Can't change channel\n", r.deviceName)
    }
}

func (r *TVRemote) PreviousChannel() {
    if r.isTurnedOn {
        r.currChannel--
        if r.currChannel < 1 {
            r.currChannel = r.maxChannel
        }
        fmt.Printf("%s channel changed to %d\n", r.deviceName, r.currChannel)
    } else {
        fmt.Printf("%s is off -> Can't change channel\n", r.deviceName)
    }
}

func (r *RadioRemote) TurnOn() {
    r.isTurnedOn = true
    fmt.Printf("%s is turned on\n", r.deviceName)
}

func (r *RadioRemote) TurnOff() {
    if r.isTurnedOn {
        r.isTurnedOn = false
        fmt.Printf("%s is off\n", r.deviceName)
    }
}

func (r *RadioRemote) SetVolume(volume int) {
    if r.isTurnedOn {
        if volume >= 0 && volume <= r.maxVolume {
            r.volume = volume
            fmt.Printf("%s volume set to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s volume out of range (0 - %d)\n", r.deviceName, r.maxVolume)
        }
    } else {
        fmt.Printf("%s is off -> Can't set volume.\n", r.deviceName)
    }
}

func (r *RadioRemote) IncreaseVolume() {
    if r.isTurnedOn {
        if r.volume < r.maxVolume {
            r.volume++
            fmt.Printf("%s volume increased to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s is already at max volume\n", r.deviceName)
        }
    } else {
        fmt.Printf("%s is off -> Can't increase volume\n", r.deviceName)
    }
}

func (r *RadioRemote) DecreaseVolume() {
    if r.isTurnedOn {
        if r.volume > 0 {
            r.volume--
            fmt.Printf("%s volume decreased to %d\n", r.deviceName, r.volume)
        } else {
            fmt.Printf("%s is already at min volume\n", r.deviceName)
        }
    } else {
        fmt.Printf("%s is off -> Can't decrease volume\n", r.deviceName)
    }
}

func (r *RadioRemote) NextChannel() {
    if r.isTurnedOn {
        r.currChannel++
        if r.currChannel > r.maxChannel {
            r.currChannel = 1
        }
        fmt.Printf("%s channel changed to %d\n", r.deviceName, r.currChannel)
    } else {
        fmt.Printf("%s is off -> Can't change channel\n", r.deviceName)
    }
}

func (r *RadioRemote) PreviousChannel() {
    if r.isTurnedOn {
        r.currChannel--
        if r.currChannel < 1 {
            r.currChannel = r.maxChannel
        }
        fmt.Printf("%s channel changed to %d\n", r.deviceName, r.currChannel)
    } else {
        fmt.Printf("%s is off -> Can't change channel\n", r.deviceName)
    }
}

func testRemote(remote TestRemote, numTests int) {
    remote.TurnOn()

    for i := 0; i < numTests; i++ {
        remote.IncreaseVolume()
        remote.IncreaseVolume()
        remote.DecreaseVolume()
        remote.DecreaseVolume()
        remote.NextChannel()
        remote.NextChannel()
        remote.PreviousChannel()
        remote.PreviousChannel()
    }

    remote.TurnOff()

    remote.DecreaseVolume()
    remote.IncreaseVolume()
    remote.NextChannel()
    remote.PreviousChannel()
}

func main() {
    numTests := 1

    tvRemote := &TVRemote{
        deviceName:  "Samsung TV",
        isTurnedOn:  false,
        volume:      49,
        currChannel: 19,
        maxVolume:   50,
        maxChannel:  20,
    }

    radioRemote := &RadioRemote{
        deviceName:  "Brondi Radio Moderna",
        isTurnedOn:  false,
        volume:      99,
        currChannel: 79,
        maxVolume:   100,
        maxChannel:  80,
    }

    fmt.Println("Testing TV Remote:")
    testRemote(tvRemote, numTests)

    fmt.Println("\nTesting Radio Remote:")
    testRemote(radioRemote, numTests)
}
