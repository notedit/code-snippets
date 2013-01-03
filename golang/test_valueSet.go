package main

import (
    "reflect"
    "fmt"
)


func main() {
    var i uint8 = 126
    v := reflect.ValueOf(i)
    fmt.Println(v.CanSet())
    fmt.Println(v.Interface().(uint8))
}
