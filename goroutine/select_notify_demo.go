package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main() {
    done := make(chan struct{})
    c := GenerateInt(done)
    wg := sync.WaitGroup{}
    time.Sleep(5 * time.Second)
    close(done)
    wg.Add(1)
    go func() {
        defer wg.Done()
        for range c {
            fmt.Println(<-c)
        }
        fmt.Println("end...")
    }()
    wg.Wait()
}

func GenerateInt(done chan struct{}) (ch chan int) {
    ch = make(chan int, 10)
    go func() {
    L:
        for {
            select {
            case ch <- rand.Int():
            case <-done:
                break L
            }
        }
        close(ch)
    }()
    return
}
