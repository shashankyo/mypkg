package main

import "fmt"

func AcceptAnything(thing interface{}) {
	fmt.Println(thing)
	whistle, ok := thing.(Whistle)
	if ok {
		whistle.MakeSound()
	}

}
func main() {
	AcceptAnything(3.145)
	AcceptAnything("A string")
	AcceptAnything(true)
	AcceptAnything(Whistle("toyco carry"))
}
func AcceptWhistle(whistle Whistle) {
	fmt.Println(whistle)
	whistle.MakeSound()
}
