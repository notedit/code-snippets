package main

import (
    "reflect"
    "fmt"
    "os"
)

type foo struct {
    fn interface{}
}

type CommonError struct {
    message string
    detail  os.Error
}

func (c CommonError)String() string {
    return c.message + c.detail.String()
}

func (f foo)Execute(param ...interface{}) {

    refv := make([]reflect.Value,len(param))

    for k,v := range param {
        refv[k] = reflect.ValueOf(v)
    }

    fmt.Printf("%#v\n",refv)
    function := reflect.ValueOf(f.fn)
    functype := reflect.TypeOf(f.fn)
    fmt.Println(functype.NumIn())
    fmt.Printf("%#v\n",functype.In(0))

    fmt.Println(functype.In(1))
    a := function.Call(refv)
    fmt.Printf("%#v\n",a)

    fmt.Println("\n",len(a))

    //if a[1].IsNil() {
    //    fmt.Printf("i do not exist %v\n\n",a[1].Kind())
    //} else {
    //    if e,ok := a[1].Interface().(os.Error); ok {
    //        fmt.Println(e.String(),"commonerror")
    //    } else {
    //        fmt.Println("oh yeah yeah yeah")
    //    }
    //}
    //fmt.Printf("%s\n",a[1].Interface().(os.Error).String())
}


func add(a int,b map[string]string ) (int,os.Error) {
    fmt.Println(a)
    fmt.Printf("%#v\n",b)
    return a,CommonError{detail:os.NewError("new error")}
    
}

func justprint(a int, b map[string]string) {
    fmt.Println(a,"\n")
    fmt.Printf("%v\n",b)
}

func main() {
    f := new(foo)
    f.fn = add
    f.fn = justprint
    a := map[string]string {"dfdf":"aaaa","ddd":"bbbb"}
    f.Execute(1,a)
}
