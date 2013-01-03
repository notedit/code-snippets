package main

import (
    "fmt"
    "bytes"
)

func main() {

    s := []byte("dfdfhelloworld")
    var o byte = 'o'
    cindex := bytes.IndexByte(s,o)
    fmt.Printf("%#v",cindex)

}
