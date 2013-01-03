package main

import (
    "fmt"
)

func main(){
    var i int16 = -123
    if i < -(1<<5)  {
        fmt.Println(i)
    }
}
