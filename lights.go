package main

import "fmt"

func lightsOn() {
	fmt.Println("Turning Lights On")
	if !isSimulator {
		pin.High()
	}

}

func lightsOff() {
	fmt.Println("Turning Lights Off")
	if !isSimulator {
		pin.Low()
	}
}
