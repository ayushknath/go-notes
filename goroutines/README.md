# Concurrency

- It is a way of handling simultaneous events
- The goal of concurrency is to structure the simultaneous events effectively

# Goroutines

- Go's implementation of concurrency
- Use the `go` keyword to make a goroutine
```go
import (
    "sync"
)

var wg sync.WaitGroup
wg.Add(number_of_goroutines)
defer wg.Done()
wg.Wait()
```

# Go channels

- A way to handle data between goroutines

```go
intChan := make(chan int)
bytesChan := make(chan []byte)

// Read from a channel
someVar :=<- intChan

// Write to a channel
intChan <- 10

// Read a channel using for loop
for nums := range intChan {
    if nums < 1 {
        break
    }
}

// Read-only channel
func readOnlyChan(ch <-chan int) {
    // something
}

// Write-only channel
func writeOnlyChan(ch chan<- int) {
    // something
}
```

- Finally close the channel using built-in `close()` function
- If not closed, memory leak occurs
- While using for loops and range to read from a channel, it will exit the loop once the channel has been closed (using built-in `close()`) or by any other mechanism (`break`)
- Channels block when they are being written to or read from
