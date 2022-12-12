package main

import (
	"fmt"
	"net/http"
)

func main() {
	checkStatus := func(done <-chan any, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					// Here we see the goroutine doing its best to signal that there's an error.
					// What else can it do? It can't pass it back. How many errors is too many?
					// Does it continue making requests?
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()

		return responses
	}

	done := make(chan any)
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response: %v\n", response.Status)
	}
}
