package main

import (
	"fmt"
	"math/rand"
)

// As a reminder, a generator for a pipeline is any function that converts a set of
// discrete values into a stream of values on a channel.

func main() {

	// This function will repeat the values you pass to it infinitely until you tell it
	// to stop.
	repeat := func(done <-chan any, values ...any) <-chan any {
		valueStream := make(chan any)
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()

		return valueStream
	}

	// This pipeline stage will only take the first num items off of its incoming valueStream and then
	// exit.
	take := func(done <-chan any, valueStream <-chan any, num int) <-chan any {
		takeStream := make(chan any)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()

		return takeStream
	}

	// This function returns an infinite channel of random integers, generated on an as-needed
	// basis.
	repeatFn := func(done <-chan any, fn func() any) <-chan any {
		valueStream := make(chan any)
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()

		return valueStream
	}

	done := make(chan any)
	defer close(done)

	random := func() any {
		return rand.Int()
	}

	// We created a repeat generator to generate an infinite number of ones, but
	// then only take the first ten. Because the repeat generator's send blocks on
	// the take stage's receive, the repeat generator is very efficient. Although we
	// have the capability of generating an infinite stream of ones, we only generate
	// N+1 instances where N is the number we pass into the take stage.
	for num := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v ", num)
	}

	for num := range take(done, repeatFn(done, random), 10) {
		fmt.Println(num)
	}
}
