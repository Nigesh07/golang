package main

import (
	"fmt"
	"time"
)

func main() {
	res := make(chan int)
	oddTomerger := make(chan int)
	evenTosquarer := make(chan int, 10)
	squarerTomerger := make(chan int)
	merged := make(chan int)
	done := make(chan struct{})

	go helper(res)
	go oddevenspliter(res, oddTomerger, evenTosquarer)
	go squarer(evenTosquarer, squarerTomerger)
	go merger(oddTomerger, squarerTomerger, merged)
	go printer(merged, done)

	<-done
}

func helper(ans chan int) {
	for i := 0; i < 10; i++ {
		ans <- i
	}
	close(ans)
}

func oddevenspliter(num chan int, odd chan int, even chan int) {
	for i := range num {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(odd)
	close(even)
}

func squarer(in chan int, out chan int) {
	for i := range in {
		time.Sleep(2 * time.Second)
		out <- i * i
	}
	close(out)
}

func merger(odd chan int, even chan int, ans chan int) {
	for odd != nil || even != nil {
		select {
		case a, ok := <-odd:
			if ok {
				ans <- a
			} else {
				odd = nil
			}
		case a, ok := <-even:
			if ok {
				ans <- a
			} else {
				even = nil
			}
		}
	}
	close(ans)
}

func printer(res chan int, done chan struct{}) {
	for i := range res {
		fmt.Println(i)
	}
	done <- struct{}{}
}
