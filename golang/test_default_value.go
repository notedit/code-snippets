package main

import (
    "fmt"
    "time"
    "sync"
)


func main(){
    type Client struct {
        Timeout     time.Duration
        lk          sync.Mutex
    }

    c := &Client{}
    fmt.Printf("%#v\n",c.Timeout)
}
