package main

import (
	"nlp_worker/mongodb"
	"sync"
)

func init() {}
func main() {
	job := make(chan string)
	var wg sync.WaitGroup
	wg.Add(1)
	go mongodb.Iter(job)
	go nlpExtractor(&wg, job)
	// job <- "bu1in07170hutlop8mf0"
	wg.Wait()
}
