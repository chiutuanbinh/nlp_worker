package main

import (
	"binhct/common/message"
	"log"
	"nlp_worker/mongodb"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func main() {
	onlineRun()
}

func offlineRun() {
	for true {
		job := make(chan string, 1)
		var wg sync.WaitGroup
		go mongodb.Iter(job)
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go nlpExtractor(&wg, job)
		}

		// job <- "bu1in07170hutlop8mf0"
		wg.Wait()
		time.Sleep(time.Second)
	}
}

func onlineRun() {
	// messagePoster := message.CreatePoster("127.0.0.1:9092", "News")
	job := make(chan string, 1)
	processFunc := func(key string, value string) {
		log.Printf("Receive mess %v %v\n", key, value)
		if key == "id" {
			job <- value
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go nlpExtractor(&wg, job)
	}

	message.Subcribe("127.0.0.1:9092", "news", processFunc)
	// job <- "bu1in07170hutlop8mf0"
	wg.Wait()
	time.Sleep(time.Second)
}
