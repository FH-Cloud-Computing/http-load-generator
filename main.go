package main

import (
	"log"
	"net/http"
	"runtime"
	"time"
)

func load(_ http.ResponseWriter, _ *http.Request) {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
	close(done)
}

func health(_ http.ResponseWriter, _ *http.Request) {
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/load", load)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to launch webserver (%v)", err)
	}
}
