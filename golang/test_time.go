
package main

import (
    "fmt"
    "time"
)

func main(){

    timeout := time.Duration(100) * time.Millisecond
    fmt.Printf("%#v\n",timeout)
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Println(now)
    }
}
