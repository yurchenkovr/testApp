package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	MaxWorker     int
	MaxQueue      int
	countOfEvents int = 1000000
)

func main() {
	fmt.Print("Enter count of workers: ")
	fmt.Scanf("%d", &MaxWorker)
	fmt.Print("Enter size of queue: ")
	fmt.Scanf("%d", &MaxQueue)

	JobQueue = make(chan Job, MaxQueue)

	limiter := time.Tick(time.Duration(float64(time.Minute) / float64(countOfEvents)))

	fmt.Println(time.Now())

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	for i := 0; i < countOfEvents; i++ {
		data := data{Key: i, Value: rand.Intn(100000)}
		job := Job{data: data}
		JobQueue <- job

		<-limiter
	}

	time.Sleep(time.Minute)
	fmt.Println(time.Now())
}
