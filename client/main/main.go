package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

	//limiter := time.Tick(time.Duration(float64(time.Minute) / float64(countOfEvents)))
	//fmt.Println("timeout: ", <-limiter)
	t := float64(60.0 / float64(countOfEvents))
	t1 := strconv.FormatFloat(t, 'f', 8, 64) + "s"
	timeout, err := (time.ParseDuration(t1))
	fmt.Println(t1)
	fmt.Println("timeout: ", timeout)
	if err != nil {
		log.Println("ERROR", err)
	}
	fmt.Println(time.Now())

	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	for i := 0; i < countOfEvents; i++ {
		data := data{Key: i, Value: rand.Intn(100000)}
		job := Job{data: data}
		JobQueue <- job

		time.Sleep(timeout)
		//<-limiter
	}

	time.Sleep(time.Minute)
	fmt.Println(time.Now())
}
