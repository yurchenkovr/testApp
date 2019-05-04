package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"testing"
)

const (
	ok    = "\u2713"
	notOk = "\u2717"
)

func BenchmarkDoJobServer(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		xdata := data{Key: i, Value: rand.Intn(100000)}
		job := Job{data: xdata}
		job.data.DoJob()
	}
}

func TestUpload(t *testing.T) {
	xdata := data{1, 100}
	url := "http://localhost:8081/"
	statusCode := 200

	content, err := json.Marshal(xdata)
	if err != nil {
		log.Println("ERROR: ", err)
	}

	t.Log("Test Uploading content")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(content))
			if err != nil {
				t.Fatal("\t\tShould be able to make the POST call - ", notOk, err)
			}
			t.Log("\t\tShould be able to make the POST call - ", ok)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status  - %s", statusCode, ok)
			} else {
				t.Errorf("\t\tShould receive a \"%d\" status - %s", statusCode, notOk)
			}
		}

	}
}
