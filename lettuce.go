// To build and distribute to lettuce.dev
// ship_lettuce

// RaspberryPi Library: go-rpio
// https://github.com/stianeikeland/go-rpio
//
// fswebcam image.jpg
// scp pi@192.168.1.126:/home/pi/image.jpg ./

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/go-rpio"
)

var pin rpio.Pin
var isSimulator bool

func main() {
	var timeEvents []timeEvent

	initTimeEvents(&timeEvents)
	fmt.Println(timeEvents)

	isSimulator = os.Getenv("LETTUCE_SIMULATOR") == "true"

	fmt.Println(isSimulator)
	if isSimulator {
		fmt.Println("Simulator Mode")
	} else {
		fmt.Println("Production Mode")
		rpio.Open()
		defer rpio.Close()
		pin = rpio.Pin(17)
		pin.Output()
	}

	// Lights are probably on when we're deploying
	lightsOn()

	lastTime := time.Now()
	for {
		fmt.Println("Tick")

		currentTime := time.Now()
		fmt.Println("Current Time:")
		fmt.Println(currentTime)

		events := findTimeEvents(timeEvents, lastTime, currentTime)
		fmt.Println("Found Events:")
		fmt.Println(events)
		fireEvents(events)

		lastTime = currentTime
		time.Sleep(1000 * time.Millisecond)
	}
}
