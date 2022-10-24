### The for-select Loop

Something you’ll see over and over again in Go programs is the for-select loop. It’s nothing more than something like this:
```go
for { // Either loop infinitely or range over something 
    select {
        // Do some work with channels
    } 
}
```
There are a couple of different scenarios where you’ll see this pattern pop up.
#### Sending iteration variables out on a channel
Oftentimes you’ll want to convert something that can be iterated over into values on a channel. This is nothing fancy, and usually looks something like this:
```go
for _, s := range []string{"a", "b", "c"} {
     select {
        case <-done: 
            return
        case stringStream <- s:
    }   
}
```
#### Looping infinitely waiting to be stopped
It’s very common to create goroutines that loop infinitely until they’re stopped. There are a couple variations of this one. Which one you choose is purely a stylis‐ tic preference.
The first variation keeps the select statement as short as possible:
```go
for { 
    select {
        case <-done: return
        default:     
    }
    // Do non-preemptable work
}
```
If the done channel isn’t closed, we’ll exit the select statement and continue on to the rest of our for loop’s body.

The second variation embeds the work in a default clause of the select statement:
```go
for { 
    select {
        case <-done: 
            return
        default:
        // Do non-preemptable work
    } 
}
```
When we enter the select statement, if the done channel hasn’t been closed, we’ll execute the default clause instead.