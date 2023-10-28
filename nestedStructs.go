package main

import (
	"fmt"
)

func main() {
    currBattery := battery{"LiOn", 5500}

    currSmartphone := smartphone{"Android 11 ZH3", "Ericsson", currBattery}
	
    fmt.Println(currSmartphone.Model)
    fmt.Println(currSmartphone.Maker)
    fmt.Println(currSmartphone.Battery, "\n")
	

	samsSmartphone := smartphone{"Samsung OS 323", "Samsung", battery{"LiPo", 4500}}

	fmt.Println(samsSmartphone.Model)
	fmt.Println(samsSmartphone.Maker)
	fmt.Println(samsSmartphone.Battery, "\n")

	}	

type smartphone struct{
	Model string
	Maker string
	Battery battery
}

type battery struct{
	Materials string
	Power int
}
