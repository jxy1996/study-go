package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

type Man struct {
    Name string `json:"name"doc:"名字"`
    Age  int    `json:"age"doc:"年龄"`
    Wife string `json:"wife,omitempty"doc:"妻子"`
}

func (m *Man) String() {
    data, _ := json.Marshal(&m)
    fmt.Println(string(data))
}

func main() {
    m := Man{
        Name: "jon",
        Age:  80,
    }
    rt := reflect.TypeOf(m)
    mName, ok := rt.FieldByName("Name")
    if ok {
        tag := mName.Tag
        fmt.Println("Tag:", tag)
        fmt.Println("json:", tag.Get("json"))
        fmt.Println("doc:", tag.Get("doc"))
    }
    
    fmt.Println("type name :", rt.Name())
    fmt.Println("type NumField :", rt.NumField())
    fmt.Println("type PkgPath :", rt.PkgPath())
    fmt.Println("type string :", rt.String())
    
    fmt.Println("type Kind.String :", rt.Kind().String())
    
    //获取结构体类型的字段名称
    for i := 0; i < rt.NumField(); i++ {
        fmt.Printf("type.Field[%d].Name:[%v]\n", i, rt.Field(i).Name)
    }
    
}
