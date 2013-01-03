

package main

import (
    "fmt"
    "time"
)

func main(){
    test_go()
}

func test_go() {
    go func(){
        fmt.Println("testdfdfdfdfdfdfd")
        time.Sleep(100 * time.Millisecond)
        fmt.Println("dfdfdfdfdfd")
    }()

    time.Sleep(101 * time.Millisecond)
    return
}
