# Remote Control Program

This Go program simulates the behavior of TV and radio remote controls. It provides a set of methods and interfaces for controlling these devices. The program allows you to interact with the remotes, turn them on and off, adjust volume, change channels, and perform various tests to observe their behavior.

## Remote Control Interfaces

The program defines two interfaces for remote controls:

1. `TestRemote`: This interface includes methods for turning the remote on and off, setting volume, increasing and decreasing volume, and changing channels. Actions are prevented if the device is turned off.

## Remote Control Devices

The program implements two types of remote control devices:

1. `TVRemote`: This struct represents a TV remote control. It includes fields for the device name, power status, volume, current channel, maximum volume, and maximum channel.

2. `RadioRemote`: This struct represents a radio remote control. It also includes fields for the device name, power status, volume, current channel, maximum volume, and maximum channel.

## Remote Control Methods

Both `TVRemote` and `RadioRemote` implement the methods defined in the `TestRemote` interface, and actions are prevented if the device is turned off:

### Common Methods
- `TurnOn()`: Turns on the remote control.
- `TurnOff()`: Turns off the remote control.

### Volume Control Methods
- `SetVolume(volume int)`: Sets the volume to the specified value.
- `IncreaseVolume()`: Increases the volume by one if not already at the maximum volume.
- `DecreaseVolume()`: Decreases the volume by one if not already at the minimum volume.

### Channel Control Methods
- `NextChannel()`: Changes the channel to the next one, wrapping around to the first channel if at the maximum channel.
- `PreviousChannel()`: Changes the channel to the previous one, wrapping around to the maximum channel if at the first channel.

## Testing

The program includes a `testRemote` function that demonstrates the functionality of the remotes. You can specify the number of test iterations to perform various actions on the remote controls and observe their behavior. The tests include turning the remote on and off, adjusting volume, and changing channels.

## Usage

To use the program, create instances of `TVRemote` and `RadioRemote`, set their properties, and call the `testRemote` function to see how the remote controls behave.

```go
tvRemote := &TVRemote{
    deviceName:  "Samsung TV",
    isTurnedOn:  false,
    volume:      10,
    currChannel: 1,
    maxVolume:   50,
    maxChannel:  20,
}

radioRemote := &RadioRemote{
    deviceName:  "Brondi Radio Moderna",
    isTurnedOn:  false,
    volume:      20,
    currChannel: 1,
    maxVolume:   100,
    maxChannel:  80,
}

numTests := 2
testRemote(tvRemote, numTests)
testRemote(radioRemote, numTests)
