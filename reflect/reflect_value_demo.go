package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

type Man1 struct {
    Name string `json:"name"doc:"名字"`
    Age  int    `json:"age"doc:"年龄"`
    Wife string `json:"wife,omitempty"doc:"妻子"`
}

func (m *Man1) String() {
    data, _ := json.Marshal(&m)
    fmt.Println(string(data))
}

func Info(i interface{}) {
    //获取value信息
    v := reflect.ValueOf(i)
    
    //通过value获取type信息
    t := v.Type()
    
    //获取类型名称：
    fmt.Println("Type :", t.Name())
    
    //访问接口字段名，字段类型和字段值信息
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i).Interface()
        
        //获取类型
        fmt.Printf("Name : [%s] , Type :  [%v] , value :[%v]\n", field.Name, field.Type, value)
    }
    
}

func main() {
    m := Man1{
        Name: "tom",
        Age:  20,
        Wife: "Lin",
    }
    m.String()
    Info(m)
}
