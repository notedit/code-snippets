package main

import (
    "fmt"
    "bytes"
    "encoding/binary"
)

func main(){
    var aa string = "\x00\x00\x00FF\x00\x00\x00\x10Operation\x00\x01\x00\x00\x00\x02Method\x00\t\x00\x00\x00Test.Add\x00\x03Argument\x00\x13\x00\x00\x00\x10A\x00\x03\x00\x00\x00\x10B\x00\x04\x00\x00\x00\x00\x00"
    buf := bytes.NewBufferString(aa)
    var msglen uint32
    err := binary.Read(buf,binary.BigEndian,&msglen)
    if err != nil {
        fmt.Printf("%v\n",err)
    }
    fmt.Printf("%v\n",msglen)
}
