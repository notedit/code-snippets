package main

import (
    "bytes"
    "fmt"
    "os"
    "encoding/gob"
    "log"
)

type P struct {
    X,Y,Z int
    Name  string
}

type Q struct {
    X,Y *int32
    Name    string
}

func main() {
    var network bytes.Buffer
    enc := gob.NewEncoder(&network)
    dec := gob.NewDecoder(&network)

    err := enc.Encode(P{3,4,5,"notedit"})
    if err != nil {
        log.Fatal("encode error",err)
    }
    fmt.Printf("%x\n",network.String())
    var q Q
    err = dec.Decode(&q)
    if err != nil {
        log.Fatal("decode error:",err)
    }
    fmt.Printf("%q: {%d,%d}\n",q.Name,*q.X,*q.Y)
   
    file_handle,err := os.OpenFile("/tmp/aaa.go",os.O_CREATE,0600)
    aa := 127
    gob.NewEncoder(file_handle).Encode(aa)
    file_handle.Close()

}
