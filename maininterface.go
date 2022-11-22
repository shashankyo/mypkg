package main

import (
	"fmt"
	"lcorequest/go/src/mypkg"
)

func main() {
	var value mypkg.myInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
