package main

import (
    "fmt"
    "launchpad.net/mgo/bson"
)

type bsonArr struct {
    Length  string
    Rowdata []byte
}

type bsonRaw struct {
    Length string
    Rowdata bson.Raw
}

func main() {

    ba := &bsonArr{Length:"123",Rowdata:[]byte{3,4,5}}

    aa ,err := bson.Marshal(ba)
    fmt.Println(err)
    fmt.Println(aa)
    
    cc := new(bsonRaw)
    _ = bson.Unmarshal(aa,cc)
    fmt.Printf("%#v\n",cc)
}
