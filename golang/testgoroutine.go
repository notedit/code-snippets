package main

import (
    "fmt"
    "time"
)

var c chan int

func ready(w string, sec int64){
    time.Sleep(sec * 1e9)
    fmt.Println(w, "is ready")
    c <- 1
}

func main() {
    c = make(chan int)
    go ready("Tee",2)
    go ready("Coffee",3)
    fmt.Println("I am waiting")
    i := 0
    L:  for {
        select {
            case <-c:
                i++
                if i > 1 {
                    break L
                }
        }
    }
}
