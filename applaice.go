package main

import "fmt"

type Appliance interface {
	TurnOn()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Heating up")
}

func (c CoffeePot) Brew() {
	fmt.Println("powering up")
}

func main() {
	var device Appliance
	device = Fan("window breeze")
	device.TurnOn()
	device = CoffeePot("luxbrew")
	device.TurnOn()
}
