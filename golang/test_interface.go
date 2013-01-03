package main

import (
    "fmt"
    "launchpad.net/gobson/bson"
)

func main() {
    var inter [4]interface{}
    in1 := 1
    in2 := [...]int{1,2,3}
    in3 := []byte("womengde")
    in4 := "dfdfdfd"


    inter[0] = in1
    inter[1] = in2
    inter[2] = in3
    inter[3] = in4

    data,_ := bson.Marshal(inter)

    fmt.Printf("%#v\n",data)
    bson.Unmarshal(data,inter)

    fmt.Printf("%#v\n",inter)


}
