package main

import (
    "errors"
    "fmt"
    "reflect"
)

type BackendError struct {
    message string
    detail  string
}

func (e BackendError) Error() string {
    return fmt.Sprintf("%s:%s",e.message,e.detail)
}


func FuncValue(a int) error {
    if a == 1 {
        a := errors.New("aadffd")
        fmt.Println(a.Error())
        return errors.New("a == 1")
    } 
    be := BackendError{message:"a != 1",detail:"a != 1111"}
    fmt.Println(be.Error())
    return be
}

func main() {
    fv := FuncValue
    fva := reflect.ValueOf(fv)
    fmt.Printf("%#v\n",reflect.ValueOf(fv))
    a := 1
    fvs := fva.Call([]reflect.Value{reflect.ValueOf(a)})
    fmt.Printf("%#v\n",fvs[0])
    fmt.Printf("%v\n",fvs[0].String())
    a = 2
    fvs = fva.Call([]reflect.Value{reflect.ValueOf(a)})
    fmt.Printf("%v\n",fvs[0].String())
    in := fvs[0].Interface()
    fmt.Printf("%v\n",in)
    switch t := in.(type) {
        case BackendError: 
            fmt.Printf("%#v\n",t)
            fmt.Println("this is a backenderror")
        case error:
            fmt.Println("this is an error")
        default:
            fmt.Println("this is default")
    }
}
