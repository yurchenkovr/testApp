package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var JobQueue chan Job

type data struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

type Job struct {
	data data
}

func (data *data) DoJob() {
	content, err := json.Marshal(data)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	makeRequest(content)
}

func makeRequest(content []byte) {
	resp, err := http.Post("http://localhost:8081/", "application/json", bytes.NewBuffer(content))
	if err != nil {
		log.Println("ERROR: ", err)
	}
	defer resp.Body.Close()
}
