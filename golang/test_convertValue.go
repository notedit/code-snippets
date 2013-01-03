package main

import (
    "database/sql/driver"
    "fmt"
    "time"
)

func main(){
    timen := time.Now()
    interList := []interface{}{"aaa","fafaf",34,timen,true}
    for _,arg := range interList {
        var err error
        a,err := driver.DefaultParameterConverter.ConvertValue(arg)
        if err != nil {
            fmt.Println(err)
        }
        fmt.Printf("%#v\n",a)
    }
}
