package main

import (
	"fmt"
	"net/http"
)

var (
	MaxWorker int
	MaxQueue  int
)

func main() {
	fmt.Print("Enter count of workers: ")
	fmt.Scanf("%d", &MaxWorker)
	fmt.Print("Enter size of queue: ")
	fmt.Scanf("%d", &MaxQueue)

	JobQueue = make(chan Job, MaxQueue)

	handler := http.NewServeMux()
	handler.HandleFunc("/", Collector)

	s := &http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	fmt.Println("Starting the dispatcher")
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
}
