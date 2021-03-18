package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var doneCh = make(chan struct{})
var exit = make(chan struct{})

var wg sync.WaitGroup

func timer(prefix string) {
	defer wg.Done()
	t := time.Now()
	for {
		select {
		case <-time.After(time.Second * 2):
			log.Printf("%s : %s\n", prefix, time.Since(t))
		case <-doneCh:
			log.Printf("%s : done, return \n", prefix)
			return
		}
	}
}

func main() {
	ss := []string{"New York", "Moon", "Paris"}
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(doneCh)
	}()

	for _, s := range ss {
		wg.Add(1)
		go timer(s)
	}

	go func() {
		wg.Wait()
		log.Println("wg finished..")
		close(exit)
	}()

	<-exit

}
