package main

import (
	"fmt"
)

func main() {
	_, err := os.open("file.text")
	fmt.Print(err)
}
