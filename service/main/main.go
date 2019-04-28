package main

import (
	"fmt"
	"net/http"
)

const (
	MaxWorker int = 2
	MaxQueue  int = 20
)

func main() {
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

	fmt.Println("END")
}
