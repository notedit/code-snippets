package main

import (
    "fmt"
)

func returntest() (b int) {
    defer func(){
        b = 3
    }()
    b = 4
    return b
}

func main(){
    fmt.Println(returntest())
}
