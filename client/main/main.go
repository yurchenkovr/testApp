package main

import (
	"math/rand"
	"time"
)

const (
	MaxWorker     int = 8
	MaxQueue      int = 1000
	countOfEvents int = 100000
)

func main() {

	JobQueue = make(chan Job, MaxQueue)

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	for i := 0; i < countOfEvents; i++ {
		data := data{key: i, value: rand.Intn(1000000)}
		job := Job{data: data}
		JobQueue <- job
	}
	time.Sleep(time.Minute)
}
