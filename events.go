package main

import "time"

type timeEvent struct {
	Action string
	Hour   int
	Minute int
	Second int
}

func fireEvents(events []timeEvent) {
	for _, event := range events {
		switch event.Action {
		case "lights_on":
			lightsOn()
		case "lights_off":
			lightsOff()
		case "moisture_reading":
			readMoistureSensor()
		}
	}
}

func initTimeEvents(events *[]timeEvent) {
	eventLightsOn := timeEvent{Action: "lights_on", Hour: 6, Minute: 0, Second: 0}
	eventLightsOff := timeEvent{Action: "lights_off", Hour: 0, Minute: 0, Second: 30}
	*events = append(*events, eventLightsOn, eventLightsOff)

	// Read moisture sensor every hour
	for hour := 0; hour < 24; hour++ {
		*events = append(*events, timeEvent{Action: "moisture_reading", Hour: hour, Minute: 0, Second: 0})
	}
}

func findTimeEvents(events []timeEvent, startTime time.Time, endTime time.Time) []timeEvent {
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
