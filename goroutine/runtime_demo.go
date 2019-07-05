package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    
    //返回当前协程数量
    fmt.Println("Goroutine num:", runtime.NumGoroutine())
    
    //设置或查询可以并发执行的协程数,参数大于1表示设置，否则为查询
    i := runtime.GOMAXPROCS(0)
    fmt.Println("GOMAXPROCS:",i)
    
    go func() {
        defer func() {
            fmt.Println("defer()")
        }()
        //Goexit()为结束当前的协程运行，会执行defer函数
        runtime.Goexit()
        
        fmt.Println("Hello world!")
    }()
    
    fmt.Println("Goroutine num:", runtime.NumGoroutine())
    time.Sleep(2 * time.Second)
}
