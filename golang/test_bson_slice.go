package main

import (
    "bson"
    "fmt"
)

func main(){
    slices := []string{"aa","bb","cc"} 
    bb,err := bson.Marshal(slices)
    if err != nil {
        fmt.Println(err.Error())
    }
    fmt.Printf("%#v\n",string(bb))
    err = bson.Unmarshal(bb,&slices)
    if err != nil {
        fmt.Printf("%v",slices)
    }
}
