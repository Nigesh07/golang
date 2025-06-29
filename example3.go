package main

import (
	"fmt"
)

func main() {
	res := make(chan int)
	squareToPrint := make(chan int)
	done := make(chan struct{})

	go helper(res)
	go squarer(res, squareToPrint)
	go print(squareToPrint, done)

	<-done
}

func helper(ans chan int) {
	for i := 0; i < 5; i++ {
		ans <- i
	}
	close(ans)
}

func squarer(in chan int, out chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func print(res chan int, done chan struct{}) {
	for i := range res {
		fmt.Println(i)
	}
	done <- struct{}{}
}
