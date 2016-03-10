package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func readTimeEventsFromStore() []timeEvent {
	filename := os.Getenv("LETTUCE_EVENT_FILENAME")
	return readTimeEventsFromCSV(filename)
}

func readTimeEventsFromCSV(filename string) []timeEvent {
	var events []timeEvent

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		events = append(events, parseCsvTimeEventRecord(record))
	}
	fmt.Println(events)
	return events
}

func parseCsvTimeEventRecord(record []string) timeEvent {
	eventHour, _ := strconv.Atoi(record[1])
	eventMinute, _ := strconv.Atoi(record[2])
	eventSecond, _ := strconv.Atoi(record[3])
	return timeEvent{Action: record[0], Hour: eventHour, Minute: eventMinute, Second: eventSecond}
}
