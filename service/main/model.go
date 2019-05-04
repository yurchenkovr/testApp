package main

import (
	"fmt"
	"sync/atomic"
)

var (
	count    uint64
	JobQueue chan Job
)

type data struct {
	Key   int `json:"key"`
	Value int `json:"value"`
}

type Job struct {
	data data
}

func (data *data) DoJob() {
	fmt.Println(data, atomic.LoadUint64(&count))
	atomic.AddUint64(&count, 1)
}
