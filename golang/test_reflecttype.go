
package main

import (
    "fmt"
    "reflect"
)

type foo struct{
    fn interface{}
}

func (f foo)Execute(){
    function := reflect.ValueOf(f.fn)
    fmt.Printf("%#v\n",function)
}

func main(){
    f := new(foo)
    f.fn = func(a int,b int){
        fmt.Println(a,b,"\n")
    }
    f.Execute()
}
