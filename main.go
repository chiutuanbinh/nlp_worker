package main

import (
	"log"
	"nlp_worker/mongodb"
	"sync"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func main() {
	offlineRun()
}

func offlineRun() {
	for true {
		job := make(chan string, 5)
		var wg sync.WaitGroup
		go mongodb.Iter(job)
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go nlpExtractor(&wg, job)
		}

		// job <- "bu1in07170hutlop8mf0"
		wg.Wait()
		time.Sleep(time.Second)
	}

}
