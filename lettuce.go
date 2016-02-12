// To build and distribute to lettuce.dev
// ship_lettuce

// RaspberryPi Library: go-rpio
// https://github.com/stianeikeland/go-rpio

package main

import (
	"fmt"
	"time"

	"github.com/go-rpio"
)

type timeEvent struct {
	Action string
	Hour   int
	Minute int
	Second int
}

var pin rpio.Pin

func main() {
	var timeEvents []timeEvent

	initTimeEvents(&timeEvents)
	fmt.Println(timeEvents)

	rpio.Open()
	defer rpio.Close()
	pin = rpio.Pin(17)
	pin.Output()

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

func fireEvents(events []timeEvent) {
	for _, event := range events {
		switch event.Action {
		case "lights_on":
			lightsOn()
		case "lights_off":
			lightsOff()
		}
	}
}

func lightsOn() {
	fmt.Println("Turning Lights On")
	pin.High()
}

func lightsOff() {
	fmt.Println("Turning Lights Off")
	pin.Low()
}

func initTimeEvents(events *[]timeEvent) {
	eventLightsOn := timeEvent{Action: "lights_on", Hour: 6, Minute: 0, Second: 0}
	eventLightsOff := timeEvent{Action: "lights_off", Hour: 0, Minute: 0, Second: 30}
	*events = append(*events, eventLightsOn, eventLightsOff)
}

func findTimeEvents(events []timeEvent, startTime time.Time, endTime time.Time) []timeEvent {
	fmt.Println("Clock Time")
	fmt.Println(startTime.Clock())

	startSecond := clockToSecond(startTime.Clock())
	endSecond := clockToSecond(endTime.Clock())

	var currentEvents []timeEvent

	for _, event := range events {
		eventSecond := clockToSecond(event.Hour, event.Minute, event.Second)

		if eventSecond < endSecond && eventSecond >= startSecond ||
			eventSecond < endSecond && endSecond < startSecond {
			currentEvents = append(currentEvents, event)
		}
	}

	return currentEvents
}

func clockToSecond(hour int, minute int, second int) int {
	return hour*60*60 + minute*60 + second
}
