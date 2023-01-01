package main

import "fmt"

// We don't know if the fact that our goroutine was canceled means the channel
// we're reading from has been canceled. For this reason, we need to wrap our
// read from the channel with a select statement that also selects from a done
// channel.
func main() {
	orDone := func(done, c <-chan any) <-chan any {
		valStream := make(chan any)
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()

		return valStream
	}

	done := make(chan any)
	defer close(done)

	myChan := make(chan any)
	defer close(myChan)

	for val := range orDone(done, myChan) {
		fmt.Println(val)
	}
}
