package main

import (
    "fmt"
)

func multi(aa ...interface{}){
    fmt.Printf("%#v",aa)
}

func main(){
    multi("fafafafafaf","fafafasfasfafaf",
            "fafafafafafafaf")
}
