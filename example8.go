package main

import (
	"log"
	"sync"
	"time"
)

var config map[string]string
var configonce sync.Once = sync.Once{}

func loadconfig() {
	time.Sleep(1 * time.Millisecond)
	log.Println("Loading configuration")
	config = map[string]string{
		"hostname": "localhost",
	}
}

func getconfig(key string) string {
	if config == nil {
		configonce.Do(loadconfig)
	}
	return config[key]
}

func dosomthing(done chan struct{}) {
	getconfig("hostname")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})
	for i := 0; i < 5; i++ {
		go dosomthing(done)
	}
	for i := 0; i < 5; i++ {
		<-done
	}
}
