package main

import (
    "fmt"
)

func main(){
    aa := make([]*int,5)

    var a *int
    var i int
    for i=1; i< 4; i++ {
        a = &i
        aa = append(aa,a)
    }

    for a 
    fmt.Printf("%v\n",aa)
}
