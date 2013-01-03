package main

import (
    "fmt"
)

func testslice(bb []byte) {
    bb[0] = 0
    bb[1] = 1
}

func testarray(bb [10]byte) {
    bb[2] = 2
    bb[3] = 3
    fmt.Printf("%v\n",bb)
}

func main(){
    aa := make([]byte,10)
    testslice(aa)
    fmt.Printf("%v\n",aa)

    var cc [10]byte 
    testarray(cc)
    fmt.Printf("%v\n",cc)
}
