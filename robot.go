package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk")
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep")
}

func (r Robot) walk() {
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func play(n NoiseMaker) {
	n.MakeSound()
	// n.walk()
}

func main() {
	play(Robot("botco ambier"))
}
