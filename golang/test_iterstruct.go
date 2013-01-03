package main

import (
    "fmt"
    "reflect"
)

type Iterstruct struct {
    Int     int
    String  string
}

func main(){
    i := Iterstruct{}
    iv := reflect.ValueOf(i)
    t := iv.Type()
    fmt.Printf("%#v\n",t)
    l := t.NumField()
    fmt.Println(l)

    keys := make([]reflect.StructField,l)

    for j := range keys {
        keys[j] = t.Field(j)
    }

    for _,key := range keys {
        mval := iv.FieldByIndex(key.Index)
        if mval.IsValid() {
            fmt.Printf("%#v\n",mval.Interface())
        }
    }
}
