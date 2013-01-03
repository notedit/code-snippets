
package main

import (
    "github.com/bmizerany/mc.go"
    "time"
    "fmt"
)

func main() {

    cn,err := mc.Dial("tcp","localhost:11211")
    if err != nil {
        panic(err)
    }

    err = cn.Set("foo","bar",0,3600,0)
    if err != nil {
        panic(err)
    }

    _,_,_ = cn.Get("foo")

    start := time.Seconds()
    for i:= 0; i< 100000;i++ {
        _,_,err = cn.Get("foo")
    }
    end := time.Seconds() - start

    fmt.Printf("%#v\n",end)

}
