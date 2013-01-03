package main

import (
    "fmt"
    "bytes"
    "launchpad.net/mgo/bson"
)

type clientResponse struct {
    messageLength   uint32
    Operation       uint8
    Reply           bson.Raw
}

func main() {
    buf := bytes.NewBufferString("'\x00\x00\x00\x10operation\x00\x02\x00\x00\x00\x03reply\x00\x0c\x00\x00\x00\x10c\x00\x0f\x00\x00\x00\x00\x00")
    aa := buf.Bytes()
    cr := clientResponse{}
    err := bson.Unmarshal(aa,&cr)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("%v\n%v\n",cr.Operation,cr.Reply)
}
