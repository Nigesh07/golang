package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name string
	age  int
}

// func (employee) name() string {
// 	return "nigesh"
// }

// type person interface {
// 	name() string
// }

func main() {
	// x := "Nigesh "
	// display(x)

	// var y person
	// y = employee{}
	// display(y)

	x := employee{"nigesh", 19}
	display(x)

}

func display(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.ValueOf(x))
}
