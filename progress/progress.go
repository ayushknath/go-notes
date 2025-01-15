package main

import (
    "fmt"
    "time"
)

func main() {
    barLen := 40
    indicator := "#"
    var interval time.Duration = 100

    for i := 0; i <= 100; i++ {
        percent := float32(i) / 100.0
        percentOfBar := int(percent * float32(barLen))

        fmt.Print("[")
        for j := 1; j <= barLen; j++ {
            if j > percentOfBar {
                fmt.Print(".")
            } else {
                fmt.Print(indicator)
            }
        }
        fmt.Printf("] %d%%\r", i)

        time.Sleep(interval * time.Millisecond)
    }

    fmt.Println()
}
