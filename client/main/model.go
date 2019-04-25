package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

var JobQueue chan Job

type data struct {
	key   int `json:"key"`
	value int `json:"value"`
}

type Job struct {
	data data
}

func (data *data) doJob() {
	content, err := json.Marshal(data)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	makeRequest(content)
}

func makeRequest(content []byte) {
	resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(content))
	if err != nil {
		log.Println("ERROR: ", err)
	}
	defer resp.Body.Close()
}
