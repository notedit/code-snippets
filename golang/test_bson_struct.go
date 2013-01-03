package main

import (
    "fmt"
    "launchpad.net/mgo/bson"
)

type Person struct {
    Name string
    Phone   string ",omitempty"
}

type OtherPerson  struct {
    Name string
    PHONEeee string
}


func main() {
    data,err := bson.Marshal(&Person{Name:"bob",Phone:"fdfdfdfd"})
    if err != nil {
        panic(err)
    }

    bb := &OtherPerson{}
    err = bson.Unmarshal(data,bb)
    fmt.Printf("%v\n",bb)
}
