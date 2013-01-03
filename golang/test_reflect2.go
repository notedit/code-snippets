package main

import (
    "fmt"
    "reflect"
)

func main(){

    sa := []string{"dfasfsfds","dfdsafas"}
    v := reflect.ValueOf(sa)
    fmt.Println(v.Kind())
    fmt.Println(v.Type())

    elet := v.Type().Elem()
    fmt.Printf("%v\n",elet)
}
