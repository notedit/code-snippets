package main

import (
    "fmt"
    "launchpad.net/mgo/bson"
    "launchpad.net/mgo"
)

type Person struct {
    Name string 
    PHONE string
}

type M bson.M

func main() {
    session,err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    c := session.DB("database").C("ids")
    result := make(M)
    err = c.Find(M{"_id":"user"}).Modify(mgo.Change{Update:M{"$inc":M{"userid":1}},New:true},result)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(result["userid"].(int))
}
