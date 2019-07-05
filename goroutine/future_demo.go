package main

import (
    "fmt"
    "time"
)

type query struct {
    sql chan string
    
    result chan string
}

func (q *query) ExecQuery() {
    go func() {
        sql := <-q.sql
        //省略数据库交互
        //获取结果
        q.result <- "result from :" + sql
    }()
}

func main() {
    q := query{make(chan string, 1), make(chan string, 1)}
    go q.ExecQuery()
    
    q.sql <- "select * from tbl_xxx"
    
    time.Sleep(1 * time.Second)
    
    fmt.Println(<-q.result)
}
