package main

import (
    "fmt"
    "sync"
)

func printNumbers(wg *sync.WaitGroup, ch <-chan int) {
    defer wg.Done()

    for i := range ch {
        fmt.Printf("Reading number: %d\n", i)
    }
}

func generateNumbers(total int, ch chan<- int) {
    for i := 1; i <= total; i++ {
        fmt.Printf("Generating number: %d\n", i)
        ch <- i
    }
}

func main() {
    var wg sync.WaitGroup
    numberChan := make(chan int)

    wg.Add(1)
    go printNumbers(&wg, numberChan)
    generateNumbers(3, numberChan)

    close(numberChan)

    fmt.Println("Waiting for goroutines to finish")
    wg.Wait()
    fmt.Println("Exiting main function")
}
