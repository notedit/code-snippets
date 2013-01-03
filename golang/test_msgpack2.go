package main

import (
    "bytes"
    "fmt"
    "github.com/ugorji/go-msgpack"
)

func main(){
    test_str := "\x82\xa1A\x07\xa1B\x00" // {'A':7,'B':0}
    dec := msgpack.NewDecoder(bytes.NewBufferString(test_str),nil)
    nilobject := &struct{}{}
    err := dec.Decode(nilobject)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(nilobject)
}
