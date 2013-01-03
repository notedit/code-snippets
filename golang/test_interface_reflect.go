package main

import (
    "fmt"
    "reflect"
    "launchpad.net/mgo/bson"
)

type Arg struct {
    a   int
    b   int
}

type Request struct {
    params interface{}
}

func main() {
    arg := &Arg{a:1,b:3}
    bb,_ := bson.Marshal(arg)
    argType := reflect.TypeOf(arg)
    req := new(Request)
    
    var out Arg

    _ = bson.Unmarshal(bb,&out)
    req.params = out

    var argv reflect.Value
    argv = reflect.New(argType.Elem())
    argv.Elem().Set(reflect.ValueOf(req.params))


    fmt.Printf("%#v\n",argv)

}
