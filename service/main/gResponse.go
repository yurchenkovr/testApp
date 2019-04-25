package main

import "fmt"

var (
	count    uint
	JobQueue chan Job
)

type gResponse struct {
	key   int `json:"key`
	value int `json:"value"`
}

type Job struct {
	gResponse gResponse
}

func (gr *gResponse) doJob() {
	fmt.Println(gr, count)
	count++
}
