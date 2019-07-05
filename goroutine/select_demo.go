package main

import (
    "fmt"
    "time"
)

func main() {
    c := make(chan string)
    
    go func() {
        c <- "str"
    }()
    
    //如果使用 time.Tick 就将其创建在循环外部，避免多次创建Tick对象
    tick := time.Tick(5 * time.Second)
    for {
        select {
        case <-time.After(6 * time.Second):
            fmt.Print("After")
        case <-tick:
            fmt.Println("Tick")
        case str := <-c:
            fmt.Println(str)
        //default:
        //    fmt.Println("default...")
        }
    }
}
