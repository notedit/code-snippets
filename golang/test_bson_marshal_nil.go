package main

import (
    "fmt"
    "launchpad.net/mgo/bson"
)

type NilStruct struct {
    Abb     string
    Nilone  interface{}
    Niltwo  interface{}
    nilthree interface{}
}

func main() {

    nils := NilStruct{Abb:"1111111",Nilone:nil,Niltwo:nil,nilthree:nil}
    bb,err := bson.Marshal(&nils)
    fmt.Printf("%v\n",bb)
    fmt.Printf("%v\n",err)
}
