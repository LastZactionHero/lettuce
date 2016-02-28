package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type moistureReading struct {
	Sensor string
	Value  uint64
}

func readMoistureSensor() {
	fmt.Println("Read Moisture Sensor")
	serialCmd := os.Getenv("LETTUCE_SERIAL_APP")

	cmd := exec.Command("python", serialCmd)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	serialReadString := out.String()

	sensorReadings := parseMoistureReading(serialReadString)
	logMoistureReadings(sensorReadings)
}

func logMoistureReadings(readings []moistureReading) {
	for _, reading := range readings {
		info.Printf("Moisture %s: %d\n", reading.Sensor, reading.Value)
	}
}

func parseMoistureReading(input string) []moistureReading {
	var readings []moistureReading

	r, _ := regexp.Compile("[AB]:[\r\n]+[0-9]+")
	readingStrs := r.FindAllString(input, -1)
	for _, readingStr := range readingStrs {
		rSensor, _ := regexp.Compile("[AB]")
		sensor := rSensor.FindString(readingStr)

		rValue, _ := regexp.Compile("[0-9]+")
		value := rValue.FindString(readingStr)

		valueInt, _ := strconv.ParseUint(value, 10, 64)
		readings = append(readings, moistureReading{Sensor: sensor, Value: valueInt})
	}

	return readings
}
