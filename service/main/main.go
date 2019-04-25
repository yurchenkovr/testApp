package main

import (
	"fmt"
	"net/http"
)

const (
	MaxWorker int = 8
	MaxQueue  int = 1000
)

func main() {
	JobQueue = make(chan Job, MaxQueue)

	handler := http.NewServeMux()
	handler.HandleFunc("/", Collector)

	s := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	fmt.Println("Starting the dispatcher")
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()

	//http.HandleFunc("/work", Collector)

	if err := s.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}
	/*if err:= http.ListenAndServe(HTTPAddr,nil); err!= nil {
		fmt.Println(err.Error())
	}*/
}
