package main

func lightsOn() {
	info.Printf("Turning Lights On")
	if !isSimulator {
		pin.High()
	}

}

func lightsOff() {
	info.Printf("Turning Lights Off")
	if !isSimulator {
		pin.Low()
	}
}
