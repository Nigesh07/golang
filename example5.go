package main

import (
	"fmt"
	"sync"
)

const val = 100000

func main() {
	nums := 0
	mutex := sync.Mutex{}
	done := make(chan struct{})
	go func() {
		for i := 0; i < val; i++ {
			mutex.Lock()
			nums++
			mutex.Unlock()
		}
		done <- struct{}{}
	}()

	go func() {

		for i := 0; i < val; i++ {
			mutex.Lock()
			nums--
			mutex.Unlock()
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	fmt.Println(nums)
}
