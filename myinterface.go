package mypkg

import "fmt"

type myInterface interface {
	MethodWithoutParameters()
	MethodWithParameters()
	MethodWithReturnValue()
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodwithoutParameters called")
}

func (m MyType) MethodWithParameters(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi methodwithreturn value"
}

func (my MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called")
}
