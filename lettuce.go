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
	"log"
	"os"
	"time"

	"github.com/go-rpio"
)

var pin rpio.Pin
var isSimulator bool
var info *log.Logger

func main() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	var timeEvents []timeEvent

	initTimeEvents(&timeEvents)

	isSimulator = os.Getenv("LETTUCE_SIMULATOR") == "true"

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
		currentTime := time.Now()

		events := findTimeEvents(timeEvents, lastTime, currentTime)
		fireEvents(events)

		lastTime = currentTime
		time.Sleep(1000 * time.Millisecond)
	}
}
