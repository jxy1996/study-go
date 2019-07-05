package main

import (
    "fmt"
    "math/rand"
)

func generateIntA(done chan struct{}) (ch chan int) {
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

func GenerateIntA(done chan struct{}) (ch chan int) {
    ch = make(chan int, 20)
    send := make(chan struct{})
    go func() {
    L:
        for {
            select {
            case ch <- <-generateIntA(send):
            case <-done:
                send <- struct{}{}
                break L
            }
        }
        close(ch)
    }()
    return
}

func main() {
    done := make(chan struct{})
    //启动生产
    ch := GenerateIntA(done)
    
    //获取生成的资源
    for i := 0; i < 100; i++ {
        fmt.Println(<-ch)
    }
    //停止生产
    done <- struct{}{}
    fmt.Println("end...")
}
