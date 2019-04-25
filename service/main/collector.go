package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Collector(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var content gResponse
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	go func(content gResponse) {
		job := Job{content}
		JobQueue <- job
	}(content)
}
