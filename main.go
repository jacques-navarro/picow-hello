package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.GP16
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})
	
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}

}
