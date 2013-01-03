package main

import (
    "fmt"
)

func main() {

    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
        fmt.Println("defer1")
        }()

    panic("dfdf")

    defer func(){
        fmt.Println("defer2")
        }()
}
