package main

import (
	"fmt"
	"time"
)

func main() {
	go helper("first")
	go helper("second")
	go helper("third")
	time.Sleep(6 * time.Second)
}

func helper(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println("%s : %d", str, i)
		time.Sleep(10 * time.Millisecond)
	}
}
