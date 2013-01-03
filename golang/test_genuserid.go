package main

import (
    "fmt"
    "math/rand"
)

func main() {
    bs := make([]byte,10)
    var i int
    for i=0;i<10;i++ {
        fmt.Println(rand.Intn(10)+48)
        bs[i] = byte(rand.Intn(10) + 48)
    }
    fmt.Printf(string(bs))
}
