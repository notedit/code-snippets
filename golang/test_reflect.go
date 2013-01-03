

package main

import (
    "fmt"
    "reflect"
)

type Error interface {
    Error() string
}

type BackendError struct{
    message string
    detail  string
}

func (e BackendError)Error() string {
    return fmt.Sprintf("%s:%s",e.message,e.detail)
}

func x(i int) error {
    if i == 1 {
        return BackendError{message:"InternalError",
                            detail:"detail"}
    } 

    return nil
}

func main() {
    be := x(1)
    bevalue := reflect.ValueOf(be)

    fmt.Printf("%#v\n",bevalue)
    fmt.Printf("%v\n",bevalue.FieldByName("message").String())
    fmt.Printf("%v\n",bevalue.FieldByName("detai").String())
    fmt.Printf("%v\n",bevalue)
    fmt.Printf("%v\n",bevalue.Type())

    bevalue = reflect.ValueOf(x(0))
    fmt.Printf("%v\n",bevalue.IsValid())
}
