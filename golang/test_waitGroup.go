package main

import (
    "fmt"
    "sync"
)

var c2 = make(chan int,4)

func main(){
    var wg sync.WaitGroup
    wg.Add(1)
    go func(){
        fmt.Printf("-----\n")
        a,b,c,d := <-c2, <-c2, <-c2, <-c2
        fmt.Println("coolIdea:", a)
        fmt.Println("coolIdea:", b)
        fmt.Println("coolIdea:", c)
        fmt.Println("coolIdea:", d)
        wg.Done()
    }()
    c2 <- 5
    c2 <- 6
    c2 <- 8
    c2 <- 2
    wg.Wait()
}
