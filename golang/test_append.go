package main

import (
    "fmt"
)

func main(){
    aa := make([]string,5)
    bb := append(aa,"testsdf")
    bb = append(bb,"dfdfdfd")
    fmt.Printf("%#v\n",aa)
    fmt.Printf("%#v\n",bb)

}
