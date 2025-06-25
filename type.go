package main

import "fmt"

type employee struct {
	name string
}

func main() {
	x := "praveen"
	print(x)
	print(employee{name: "nigesh"})
}

func print(data interface{}) {

	switch val := data.(type) {
	case string:
		fmt.Printf("%s string\n ", val)
	case int:
		fmt.Printf("%d\n integer\n ", val)
	case employee:
		fmt.Printf("%s\n employee", val.name)
	default:
		fmt.Printf("\ninvalid")
	}
}
