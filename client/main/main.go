package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxWorker     int = 2
	MaxQueue      int = 20
	countOfEvents int = 1000000
)

func main() {

	JobQueue = make(chan Job, MaxQueue)

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	for i := 0; i < countOfEvents; i++ {
		data := data{Key: i, Value: rand.Intn(1000000)}
		job := Job{data: data}
		JobQueue <- job
	}
	time.Sleep(time.Minute)
	fmt.Println("Done")
}
