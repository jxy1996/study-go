package main

import (
    "fmt"
    "net/http"
    "sync"
)

/**
   main goroutine通过调用 add() 设置需要等待goroutine的数量，每一个
goroutine结束时调用Done() ,Wait()被main用来等待所有goroutine完成。
*/

var wg sync.WaitGroup

func main() {
    urns := []string{
        "https://www.baidu.com/",
        "https://www.qq.com/",
    }
    for i := range urns {
        wg.Add(1)
        getStatus(urns[i])
    }
    wg.Wait()
}

func getStatus(urn string) {
    defer wg.Done()
    resp, err := http.Get(urn)
    if err != nil {
        panic(err)
    }
    fmt.Println(resp.Status)
}
